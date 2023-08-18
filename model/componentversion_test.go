package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestComponentVersion(t *testing.T) {

	jsonObj := []byte(`{
		"name": "Hello World;v1.0.0",
		"domain": {
			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": ""
	  }`)

	expected := `{"domain":{"_key":"bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle"},"name":"Hello World;v1.0.0","objtype":"ComponentVersion"}`

	var compver2nft ComponentVersion // define user object to marshal into

	json.Unmarshal(jsonObj, &compver2nft) // convert json string into the user object

	if byteValue, err := json.Marshal(compver2nft); err == nil {
		fmt.Printf("%s\n", string(byteValue))
		fmt.Printf("%s\n", expected)
	}

}

/*
nftJSON := compver2nft.MarshalNFT(cid2json) // generate the cid and nft json for user object
// fmt.Printf("%s=%s\n", compver2nft.Key, compver2nft.NftJSON)
assert.Equal(t, nftJSON, expected, "check nft json against expected results")

var nft2compver ComponentVersion // define user object to marshal into

nft2compver.Key = compver2nft.Key         // set the nft json
nft2compver.UnmarshalNFT(cid2json)        // convert the json string into the user object
check := nft2compver.MarshalNFT(cid2json) // recalcuate the cid and nft json for the new user object
assert.Equal(t, check, expected, "check unmarshalled user against expected results") */
