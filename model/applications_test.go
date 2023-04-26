package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplications(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreih7gkkawdnx32lxlegitjtwszo35zz75sb7c643obgc3jgfmiijsy",
		"applications": [{
			"_key": "bafkreia4ioz2a6o3w5ijarqbxwfixcmevlqukjd4bndkw4bj7vosrjqfh4",
			"name": "Hello App;v1",
			"domain": {
			  "_key": "bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle",
			  "name": "GLOBAL.My Project"
			},
			"parent_key": "",
			"predecessor_key": "",
			"deployments": [121]
		  }
		]
	}`)

	expected := `{"applications":[{"_key":"bafkreia4ioz2a6o3w5ijarqbxwfixcmevlqukjd4bndkw4bj7vosrjqfh4"}]}`

	var apps2nft Applications // define user object to marshal into

	json.Unmarshal(jsonObj, &apps2nft) // convert json string into the user object
	apps2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", apps2nft.Key, apps2nft.NftJSON)
	assert.Equal(t, apps2nft.NftJSON, expected, "check nft json against expected results")

	var nft2apps Applications // define user object to marshal into

	nft2apps.NftJSON = expected     // set the nft json
	nft2apps.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2apps.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, nft2apps.NftJSON, expected, "check unmarshalled user against expected results")

}
