// Package database - Handles all interaction with ArangoDB and Long Term Storage (LTS).
// Contains utility functions for marshaling/unmarshaling json to cid/nfts
package database

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"time"

	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"github.com/arangodb/go-driver/v2/connection"
	"github.com/cenkalti/backoff"
	cid "github.com/ipfs/go-cid"
	"github.com/sanity-io/litter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	mc "github.com/multiformats/go-multicodec"

	mh "github.com/multiformats/go-multihash"
)

//lint:file-ignore S1034 Ignore all assignments for switch statements

var logger = InitLogger() // setup the logger

// DBConnection is the structure that defined the database engine and collections
type DBConnection struct {
	Collections map[string]arangodb.Collection
	Database    arangodb.Database
}

// Define a struct to hold the index definition
type indexConfig struct {
	Collection string
	IdxName    string
	IdxField   string
}

// Define a struct to hold the graph definition
type graphConfig struct {
	GraphName      string
	Collection     string
	EdgeCollection string
	From           string
	To             string
}

var initDone = false          // has the data been initialized
var dbConnection DBConnection // database connection definition

// GetEnvDefault is a convenience function for handling env vars
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key) // get the env var
	if !ex {                     // not found return default
		return defVal
	}
	return val // return value for env var
}

// InitLogger sets up the Zap Logger to log to the console in a human readable format
func InitLogger() *zap.Logger {
	prodConfig := zap.NewProductionConfig()
	prodConfig.Encoding = "console"
	prodConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	prodConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	logger, _ := prodConfig.Build()
	return logger
}

func dbJSONHTTPConnectionConfig(endpoint connection.Endpoint, dbuser string, dbpass string) connection.HttpConfiguration {
	return connection.HttpConfiguration{
		Authentication: connection.NewBasicAuth(dbuser, dbpass),
		Endpoint:       endpoint,
		ContentType:    connection.ApplicationJSON,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // #nosec G402
			},
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 90 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

// InitializeDatabase is the function for connecting to the db engine, creating the database and collections
func InitializeDatabase() DBConnection {

	const initialInterval = 10 * time.Second
	const maxInterval = 2 * time.Minute

	var db arangodb.Database
	var collections map[string]arangodb.Collection
	const databaseName = "ortelius"

	ctx := context.Background()

	if initDone {
		return dbConnection
	}

	False := false
	dbhost := GetEnvDefault("ARANGO_HOST", "localhost")
	dbport := GetEnvDefault("ARANGO_PORT", "8529")
	dbuser := GetEnvDefault("ARANGO_USER", "root")
	dbpass := GetEnvDefault("ARANGO_PASS", "")
	dburl := GetEnvDefault("ARANGO_URL", "http://"+dbhost+":"+dbport)

	var client arangodb.Client

	//
	// Database connection with backuoff retry
	//

	// Configure exponential backoff
	bo := backoff.NewExponentialBackOff()
	bo.InitialInterval = initialInterval
	bo.MaxInterval = maxInterval
	bo.MaxElapsedTime = 0 // Set to 0 for indefinite retries

	// Retry logic
	err := backoff.RetryNotify(func() error {
		fmt.Println("Attempting to connect to ArangoDB")
		endpoint := connection.NewRoundRobinEndpoints([]string{dburl})
		conn := connection.NewHttpConnection(dbJSONHTTPConnectionConfig(endpoint, dbuser, dbpass))

		client = arangodb.NewClient(conn)

		// Ask the version of the server
		versionInfo, err := client.Version(context.Background())
		if err != nil {
			return err
		}

		logger.Sugar().Infof("Database has version '%s' and license '%s'\n", versionInfo.Version, versionInfo.License)
		return nil

	}, bo, func(err error, _ time.Duration) {
		// Optionally, you can add a message here to be printed after each retry
		fmt.Printf("Retrying connection to ArangoDB: %v\n", err)
	})

	if err != nil {
		logger.Sugar().Fatalf("Backoff Error %v\n", err)
	}

	//
	// Database creation
	//

	exists := false
	dblist, _ := client.Databases(ctx)

	for _, dbinfo := range dblist {
		if dbinfo.Name() == databaseName {
			exists = true
			break
		}
	}

	if exists {
		var options arangodb.GetDatabaseOptions
		if db, err = client.GetDatabase(ctx, databaseName, &options); err != nil {
			logger.Sugar().Fatalf("Failed to create Database: %v", err)
		}
	} else {
		if db, err = client.CreateDatabase(ctx, databaseName, nil); err != nil {
			logger.Sugar().Fatalf("Failed to create Database: %v", err)
		}
	}

	//
	// Collection creation for document storage
	//

	collections = make(map[string]arangodb.Collection)
	collectionNames := []string{"applications", "components", "sbom", "vulns", "purls", "readmes", "licenses", "swagger"}

	for _, collectionName := range collectionNames {
		var col arangodb.Collection

		exists, _ = db.CollectionExists(ctx, collectionName)
		if exists {
			var options arangodb.GetCollectionOptions
			if col, err = db.GetCollection(ctx, collectionName, &options); err != nil {
				logger.Sugar().Fatalf("Failed to use collection: %v", err)
			}
		} else {
			if col, err = db.CreateCollectionV2(ctx, collectionName, nil); err != nil {
				logger.Sugar().Fatalf("Failed to create collection: %v", err)
			}
		}

		collections[collectionName] = col
	}

	collectionNames = []string{"purl2vulns", "comp2readmes", "comp2licenses", "comp2swagger"}
	for _, edgeCollectionName := range collectionNames {
		var col arangodb.Collection

		// Check if the edge collection exists
		var options arangodb.GetCollectionOptions
		col, err = db.GetCollection(ctx, edgeCollectionName, &options)
		if shared.IsNotFound(err) {
			edgeType := arangodb.CollectionTypeEdge
			col, err = db.CreateCollectionV2(ctx, edgeCollectionName, &arangodb.CreateCollectionPropertiesV2{
				Type: &edgeType,
			})
			if err != nil {
				logger.Sugar().Fatalf("Failed to create edge collection: %v", err)
			}
			logger.Sugar().Infoln("Edge collection created.")
		} else if err != nil {
			logger.Sugar().Fatalf("Failed to get edge collection: %v", err)
		}

		collections[edgeCollectionName] = col
	}

	//
	// Index creation for each collection.
	//

	idxList := []indexConfig{
		{Collection: "vulns", IdxName: "package_name", IdxField: "affected[*].package.name"},
		{Collection: "vulns", IdxName: "package_purl", IdxField: "affected[*].package.purl"},
		{Collection: "sbom", IdxName: "sbom_cid", IdxField: "cid"},
		{Collection: "purls", IdxName: "purls_idx", IdxField: "purl"},
	}

	for _, idx := range idxList {

		found := false

		if indexes, err := collections[idx.Collection].Indexes(ctx); err == nil {
			for _, index := range indexes {
				if idx.IdxName == index.Name {
					found = true
					break
				}
			}
		}

		if !found {
			// Define the index options
			indexOptions := arangodb.CreatePersistentIndexOptions{
				Unique: &False,
				Sparse: &False,
				Name:   idx.IdxName,
			}

			// Create the index
			_, _, err = collections[idx.Collection].EnsurePersistentIndex(ctx, []string{idx.IdxField}, &indexOptions)
			if err != nil {
				logger.Sugar().Fatalln("Error creating index:", err)
			}
		}
	}

	//
	// Graph creation for managing relationships
	//

	graphList := []graphConfig{
		{GraphName: "vulnGraph", Collection: "purl2vulns", From: "purls", To: "vulns"},               // one purl to many vulns
		{GraphName: "readmeGraph", Collection: "comp2readmes", From: "components", To: "readmes"},    // many comps to one readme
		{GraphName: "licenseGraph", Collection: "comp2licenses", From: "components", To: "licenses"}, // many comps to one license
		{GraphName: "swaggerGraph", Collection: "comp2swagger", From: "components", To: "swagger"},   // many comps to one swagger
	}

	for _, grph := range graphList {
		var graphOpts arangodb.GetGraphOptions

		_, err = db.Graph(ctx, grph.GraphName, &graphOpts)

		if shared.IsNotFound(err) {
			// Graph does not exist, create it
			logger.Sugar().Infof("Graph does not exist. Creating... %v", err)

			// Define the edge definitions and vertex collections
			edgeDefinition := arangodb.EdgeDefinition{
				Collection: grph.Collection,
				From:       []string{grph.From},
				To:         []string{grph.To},
			}

			def := arangodb.GraphDefinition{
				EdgeDefinitions: []arangodb.EdgeDefinition{edgeDefinition},
			}

			var options arangodb.CreateGraphOptions
			_, err = db.CreateGraph(ctx, grph.GraphName, &def, &options)
			if err != nil {
				logger.Sugar().Fatalf("Failed to create graph: %v", err)
			}
			logger.Sugar().Infoln("Graph created.")
		}
	}

	initDone = true

	dbConnection = DBConnection{
		Database:    db,
		Collections: collections,
	}

	return dbConnection
}

// PersistOnLTS interacts with the db abstraction microservice to
// store the cid/json data on NFT Storage or the OCI registry
func PersistOnLTS(cid2json map[string]string) {

	logger.Sugar().Infof("%+v\n", cid2json)
}

// FetchFromLTS interacts with the db abstraction microservice to
// fetch the json from NFT Storage or OCI registry
func FetchFromLTS(key string) (string, map[string]string) {

	msg := `{"objtype":"Domain","name":"GLOBAL"}`

	cid2json := make(map[string]string, 1)
	cid2json[key] = msg
	return key, cid2json
}

// flattenData recursively flattens a JSON string using periods to separate the nested keys. Arrays are represented by 4 digit keys
func flattenData(y interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	var flatten func(x interface{}, name string)
	flatten = func(x interface{}, name string) {
		switch v := x.(type) {
		case map[string]interface{}:
			for a, b := range v {
				flatten(b, name+a+".")
			}
		case []interface{}:
			for i, a := range v {
				flatten(a, name+fmt.Sprintf("%04d.", i))
			}
		default:
			out[name[:len(name)-1]] = x
		}
	}

	flatten(y, "")
	return out
}

// getCid takes the cid and retrieves the corresponding JSON str from LTS
func getCid(cid string) (string, bool) {

	filename := "nfts/" + cid + ".nft"

	if _, err := os.Stat(filename); err == nil {
		cidFile, _ := os.Open(filename)
		byteValue, _ := io.ReadAll(cidFile)
		cidFile.Close()

		return string(byteValue), true
	}
	return "", false
}

// genCid takes a JSON string and calculates the corresponding immutable IPFS CID
func genCid(jsonStr string) string {
	var pref = cid.Prefix{
		Version:  1,
		Codec:    uint64(mc.Raw),
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}

	_cid, err := pref.Sum([]byte(jsonStr))

	if err != nil {
		return ""
	}

	return _cid.String()
}

// splitJSON will split the json string on the bellow characters in order to just have keys and values
func splitJSON(r rune) bool {
	return r == ':' || r == ',' || r == '"' || r == '{' || r == '}' || r == '[' || r == ']'
}

// addKey2Obj will add Key=cid for nested objects
func addKey2Obj(obj any, group string, cid string) {
	fname := cases.Title(language.Und, cases.NoLower).String(group) // change group to match struct field name
	v := reflect.ValueOf(obj)

	// Check if v is a pointer, and if so, get the underlying value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Now v represents the struct, and you can call FieldByName safely
	f := v.FieldByName(fname)

	if f.IsValid() && f.CanSet() && // found the field in the object and make sure we can update the field value
		reflect.TypeOf(f.Interface()).Kind() != reflect.Array && reflect.TypeOf(f.Interface()).Kind() != reflect.Slice { // make sure its not an array/slice
		fkey := reflect.ValueOf(f.Interface()).Elem().FieldByName("Key") // see of the object we found contains the Key field
		if fkey.IsValid() && fkey.CanSet() {                             // found the field in the object and make sure we can update the field value
			fkey.SetString(cid) // set Key=cid in the object
		}
	} else if strings.Count(fname, ".") > 0 { // make sure we are working with a nested array/slice
		parts := strings.Split(fname, ".")           // split the name so we can get the object name and array index
		key := parts[len(parts)-2]                   // get the object field name
		idx := parts[len(parts)-1]                   // get the index of the array we are working with
		if i, err := strconv.Atoi(idx); err == nil { // make sure its a valid array index
			f := reflect.ValueOf(obj).Elem().FieldByName(key) // get the field for the array/slice
			if f.IsValid() && f.CanSet() {                    // found the field in the object and make sure we can update the field value
				fidx := f.Index(i).Interface()                          // get the object using the index from the array/slice
				fkey := reflect.ValueOf(fidx).Elem().FieldByName("Key") // see of the object we found contains the Key field
				if fkey.IsValid() && fkey.CanSet() {                    // found the field in the object and make sure we can update the field value
					fkey.SetString(cid) // set Key=cid in the object
				}
			}
		}
	}
}

// MakeNFT normalizes the object into the corresponding cids=json string handling nested objects
// Parameters: object of any type
// Returns: CID for the object and json string suitable for storing in Arango (ie. _key=cid and objtype added)
func MakeNFT(obj any) (string, string) {

	jsonStr := ""

	if byteValue, err := json.Marshal(obj); err == nil {
		jsonStr = string(byteValue)
	}

	/* 	objtype := reflect.TypeOf(obj).String()

	   	if strings.Count(objtype, ".") > 0 {
	   		parts := strings.Split(objtype, ".")
	   		objtype = parts[len(parts)-1]
	   	} */
	repl := regexp.MustCompile(`"_key":\s*".*","`)
	repl.ReplaceAllString(jsonStr, "")

	rootCid := ""
	jsonMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		fmt.Println(err)
		return "", ""
	}

	out := flattenData(jsonMap)

	cidmap := make(map[string]string) // output dict of grouping to json

	for len(out) > 0 {
		keys := make([]string, 0, len(out))
		groupmap := make(map[string][]string)

		for k := range out {
			keys = append(keys, k)
		}

		// sort the keys longest (most dots) and then by alpha
		sort.SliceStable(keys, func(i, j int) bool {
			lcnt := strings.Count(keys[i], ".")
			rcnt := strings.Count(keys[j], ".")

			if lcnt == rcnt {
				return (strings.Compare(keys[i], keys[j]) < 0)
			}
			return lcnt > rcnt
		})

		// find first grouping
		saveGrp := ""
		for _, k := range keys {
			parts := strings.Split(k, ".")
			key := ""
			currentGrp := ""

			if len(parts) > 1 {
				key = parts[len(parts)-1]
				currentGrp = strings.Join(parts[:len(parts)-1], ".")
			} else if len(parts) == 1 {
				currentGrp = "root"
				key = parts[0]
			}

			if currentGrp != saveGrp && saveGrp != "" {
				break
			}
			saveGrp = currentGrp

			jstr := ""

			if _, err := strconv.Atoi(key); err == nil {
				if _, ok := (out[k]).(string); ok {
					jstr = fmt.Sprintf("\"%s\"", out[k].(string))
				} else {
					jstr = fmt.Sprintf("%v", out[k])
				}
			} else {
				if _, ok := (out[k]).(string); ok {
					jstr = fmt.Sprintf("\"%s\":\"%s\"", key, out[k].(string))
				} else {
					jstr = fmt.Sprintf("\"%s\": %v", key, out[k])
				}
			}

			if jlist, ok := groupmap[currentGrp]; ok {
				groupmap[currentGrp] = append(jlist, jstr)
			} else {
				jlist := []string{jstr}
				groupmap[currentGrp] = jlist
			}
			delete(out, k)
		}

		for group := range groupmap {
			jsonStrings := groupmap[group]

			primitiveCnt := 0

			// If only one entry in the jsonStrings and its a primitive then its a primitive
			// If more then one entry in the jsonStrings and they are all privitives then its an array
			// otherwise its an array or a map

			for _, jsonString := range jsonStrings {
				var result interface{}

				err := json.Unmarshal([]byte(jsonString), &result)
				if err != nil {
					continue
				}

				switch result.(type) {
				case []interface{}:
					break
				case map[string]interface{}:
					break
				default:
					primitiveCnt++
				}
			}

			jsonStr := ""
			if primitiveCnt == len(jsonStrings) && primitiveCnt > 1 { // array of primitives, dont sort
				jsonStr = "[" + strings.Join(jsonStrings, ",") + "]"
			} else {
				sortedJSON := jsonStrings
				sort.Strings(sortedJSON)
				jsonStr = "{" + strings.Join(sortedJSON, ",") + "}"
			}

			cid := genCid(jsonStr)
			cidmap[cid] = jsonStr

			if group != "root" {
				out[group] = cid // group = nested struct path

				addKey2Obj(obj, group, cid) // Add Key=cid for nested objects
			} else {
				rootCid = cid
			}

			if err := os.Mkdir("nfts", 0755); err != nil && !os.IsExist(err) {
				fmt.Println(err)
			}
			if err := os.WriteFile("nfts/"+cid+".nft", []byte(jsonStr), 0600); err != nil {
				fmt.Println(err)
			}
		}
	}

	f := reflect.ValueOf(obj).Elem().FieldByName("Key")
	if f.IsValid() && f.CanSet() {
		f.SetString(rootCid)
	}

	dbStr := fmt.Sprintf("{\"_key\":\"%s\",%s", rootCid, jsonStr[1:])
	return rootCid, dbStr
}

// MakeJSON converts a CID back into a json string.  It will resolve any nested cids and expand as well.
// Parameters: CID to expand
// Returns: expanded json string and bool if the cid exists or not
func MakeJSON(cid string) (string, bool) {
	jsonStr, exists := getCid(cid)

	if exists {
		for {
			parts := strings.FieldsFunc(jsonStr, splitJSON)

			replaceCnt := 0
			for _, k := range parts {

				if len(k) != 59 { // skip keys that are not the length of a cid
					continue
				}

				if jsonPart, found := getCid(k); found {
					// need regex to handle replace in one go
					jsonStr = strings.ReplaceAll(jsonStr, "\""+k+"\"", jsonPart)
					jsonStr = strings.ReplaceAll(jsonStr, "["+k+"]", "["+jsonPart+"]")
					jsonStr = strings.ReplaceAll(jsonStr, "["+k+",", "["+jsonPart+",")
					jsonStr = strings.ReplaceAll(jsonStr, ","+k+",", ","+jsonPart+",")
					jsonStr = strings.ReplaceAll(jsonStr, ","+k+"]", ","+jsonPart+"]")
					replaceCnt++
				}
			}

			if replaceCnt == 0 {
				break
			}
		}
	}

	return jsonStr, exists
}

// EmptyJSON will convert a struct into an empty JSON string that includes all fields and nested fields.
func EmptyJSON(obj any) string {
	structStr := litter.Sdump(obj)

	r := regexp.MustCompile("&.*{")
	structStr = r.ReplaceAllString(structStr, "{")

	r = regexp.MustCompile("time.Time{}")
	structStr = r.ReplaceAllString(structStr, "\"\"")

	r = regexp.MustCompile("nil")
	structStr = r.ReplaceAllString(structStr, "[]")

	r = regexp.MustCompile(`([^\s]+):`)
	structStr = r.ReplaceAllString(structStr, "\"$1\":")

	r = regexp.MustCompile(`,\n\s*}`)
	structStr = r.ReplaceAllString(structStr, "\n}")

	r = regexp.MustCompile(`"Key"`)
	structStr = r.ReplaceAllString(structStr, "\"_key\"")

	structStr = strings.ToLower(structStr)

	dst := &bytes.Buffer{}
	if err := json.Compact(dst, []byte(structStr)); err != nil {
		fmt.Printf("%+v", err)
		return ""
	}
	return dst.String()
}
