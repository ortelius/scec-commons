// Package model - NFT defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

//lint:file-ignore S1034 Ignore all assignments for switch statements

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	cid "github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
)

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

	rootCid := ""
	jsonMap := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &jsonMap)
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
				jstr = fmt.Sprint(out[k])
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

	objtype := reflect.TypeOf(obj).String()

	if strings.Count(objtype, ".") > 0 {
		parts := strings.Split(objtype, ".")
		objtype = parts[len(parts)-1]
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
