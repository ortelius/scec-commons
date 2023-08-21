// Package database - Handles all interaction with ArangoDB and Long Term Storage (LTS).
// Contains utility functions for marshalling/unmaeshalling json to cid/nfts
package database

import (
	"context"
	"os"
	"reflect"
	"regexp"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	cid "github.com/ipfs/go-cid"

	mc "github.com/multiformats/go-multicodec"

	mh "github.com/multiformats/go-multihash"
)

//lint:file-ignore S1034 Ignore all assignments for switch statements

var logger = InitLogger() // setup the logger

// DBConnection is the structure that defined the database engine and collections
type DBConnection struct {
	Collection driver.Collection
	Database   driver.Database
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

// InitializeDB is the function for connecting to the db engine, creating the database and collections
func InitializeDB() DBConnection {

	var db driver.Database
	var col driver.Collection
	var conn driver.Connection
	var client driver.Client
	var err error
	const databaseName = "ortelius"

	ctx := context.Background()

	if initDone {
		return dbConnection
	}

	dbhost := GetEnvDefault("ARANGO_HOST", "localhost")
	dbport := GetEnvDefault("ARANGO_PORT", "8529")
	dbuser := GetEnvDefault("ARANGO_USER", "root")
	dbpass := GetEnvDefault("ARANGO_PASS", "")
	dburl := GetEnvDefault("ARANGO_URL", "http://"+dbhost+":"+dbport)

	if conn, err = http.NewConnection(http.ConnectionConfig{Endpoints: []string{dburl}}); err != nil {
		logger.Sugar().Fatalf("Failed to create HTTP connection: %v", err)
	}

	_, err = conn.SetAuthentication(driver.BasicAuthentication(dbuser, dbpass))

	if err == nil {
		if client, err = driver.NewClient(driver.ClientConfig{Connection: conn}); err != nil {
			logger.Sugar().Fatalf("Failed to create Client: %v", err)
		}

		exists := false
		dblist, _ := client.Databases(ctx)

		for _, dbinfo := range dblist {
			if dbinfo.Name() == databaseName {
				exists = true
				break
			}
		}

		if exists {
			if db, err = client.Database(ctx, databaseName); err != nil {
				logger.Sugar().Fatalf("Failed to create Database: %v", err)
			}
		} else {
			if db, err = client.CreateDatabase(ctx, databaseName, nil); err != nil {
				logger.Sugar().Fatalf("Failed to create Database: %v", err)
			}
		}

		exists, _ = db.CollectionExists(ctx, "evidence")
		if exists {
			if col, err = db.Collection(ctx, "evidence"); err != nil {
				logger.Sugar().Fatalf("Failed to use collection: %v", err)
			}
		} else {
			if col, err = db.CreateCollection(ctx, "evidence", nil); err != nil {
				logger.Sugar().Fatalf("Failed to create collection: %v", err)
			}
		}

		initDone = true

		dbConnection = DBConnection{
			Database:   db,
			Collection: col,
		}
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
				flatten(a, name+fmt.Sprintf("%03d.", i))
			}
		default:
			out[name[:len(name)-1]] = x
		}
	}

	flatten(y, "")
	return out
}

func getCid(cid string) (string, bool) {

	filename := "nfts/" + cid + ".nft"

	if _, err := os.Stat(filename); err == nil {
		cidFile, _ := os.Open(filename)
		byteValue, _ := ioutil.ReadAll(cidFile)
		cidFile.Close()

		return string(byteValue), true
	}
	return "", false
}

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

func splitJSON(r rune) bool {
	return r == ':' || r == ',' || r == '"' || r == '{' || r == '}' || r == '[' || r == ']'
}

// MakeNFT normalizes the object into the corresponding cids=json string handling nested objects
// Parameters: object of any type
// Returns: CID for the object and json string suitable for storing in Arango (ie. _key=cid and objtype added)
func MakeNFT(obj any) (string, string) {

	jsonStr := ""

	if byteValue, err := json.Marshal(obj); err == nil {
		jsonStr = string(byteValue)
	}

	objtype := reflect.TypeOf(obj).String()

	if strings.Count(objtype, ".") > 0 {
		parts := strings.Split(objtype, ".")
		objtype = parts[len(parts)-1]
	}
	repl := regexp.MustCompile(`"_key":\s*".*","`)
	repl.ReplaceAllString(jsonStr, "")

	rootCid := ""
	jsonMap := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &jsonMap)
	out := flattenData(jsonMap)
	out["objtype"] = objtype

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
			sortedJSON := groupmap[group]
			sort.Strings(sortedJSON)

			jsonStr := ""
			if strings.Contains(strings.Join(sortedJSON, ","), ":") {
				jsonStr = "{" + strings.Join(sortedJSON, ",") + "}"
			} else {
				jsonStr = "[" + strings.Join(sortedJSON, ",") + "]"
			}

			cid := genCid(jsonStr)
			cidmap[cid] = jsonStr

			if group != "root" {
				out[group] = cid
			} else {
				rootCid = cid
			}

			if err := os.Mkdir("nfts", 0755); err != nil && !os.IsExist(err) {
				fmt.Println(err)
			}
			os.WriteFile("nfts/"+cid+".nft", []byte(jsonStr), 0644)
		}
	}

	f := reflect.ValueOf(obj).Elem().FieldByName("Key")
	if f.IsValid() && f.CanSet() {
		f.SetString(rootCid)
	}

	dbStr := fmt.Sprintf("{\"_key\":\"%s\",\"objtype\":\"%s\",%s", rootCid, objtype, jsonStr[1:])
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

				if len(k) != 59 { // skip keys that are not the lenght of a cid
					continue
				}

				if jsonPart, found := getCid(k); found {
					// need regex to handle replace in one go
					jsonStr = strings.Replace(jsonStr, "\""+k+"\"", jsonPart, -1)
					jsonStr = strings.Replace(jsonStr, "["+k+"]", "["+jsonPart+"]", -1)
					jsonStr = strings.Replace(jsonStr, "["+k+",", "["+jsonPart+",", -1)
					jsonStr = strings.Replace(jsonStr, ","+k+",", ","+jsonPart+",", -1)
					jsonStr = strings.Replace(jsonStr, ","+k+"]", ","+jsonPart+"]", -1)
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
