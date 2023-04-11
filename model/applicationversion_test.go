package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationVersion(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreia4ioz2a6o3w5ijarqbxwfixcmevlqukjd4bndkw4bj7vosrjqfh4",
		"name": "Hello App;v1",
		"domain": {
		  "_key": "bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle",
		  "name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": "",
		"deployments": [121]
	  }`)

	expected := `{"deployments":[121],"domain":{"_key":"bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle"},"name":"Hello App;v1","objtype":"ApplicationVersion"}`

	var appver2nft ApplicationVersion // define user object to marshal into

	json.Unmarshal(jsonObj, &appver2nft) // convert json string into the user object
	appver2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", appver2nft.Key, appver2nft.NftJSON)
	assert.Equal(t, appver2nft.NftJSON, expected, "check nft json against expected results")

	var nft2appver ApplicationVersion // define user object to marshal into

	nft2appver.NftJSON = expected     // set the nft json
	nft2appver.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2appver.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, nft2appver.NftJSON, expected, "check unmarshalled user against expected results")

}
