package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponentVersion(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreieu66waq6jcefgbaxlwkeg6cnqoj5zlc63wghddh3ngtzh7olp37u",
		"name": "Hello World;v1.0.0",
		"domain": {
			"_key": "bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle",
			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": ""
	  }`)

	expected := `{"domain":{"_key":"bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle"},"name":"Hello World;v1.0.0","objtype":"ComponentVersion"}`

	var compver2nft ComponentVersion // define user object to marshal into

	json.Unmarshal(jsonObj, &compver2nft) // convert json string into the user object
	compver2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", compver2nft.Key, compver2nft.NftJSON)
	assert.Equal(t, compver2nft.NftJSON, expected, "check nft json against expected results")

	var nft2compver ComponentVersion // define user object to marshal into

	nft2compver.NftJSON = expected     // set the nft json
	nft2compver.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2compver.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, nft2compver.NftJSON, expected, "check unmarshalled user against expected results")

}
