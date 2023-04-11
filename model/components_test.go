package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponents(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreibpknufg2ciqkiqlupkglpmlrck5askj3sscieiedwwtwrcffdzwe",
		"components": [{
				"_key": "bafkreieu66waq6jcefgbaxlwkeg6cnqoj5zlc63wghddh3ngtzh7olp37u",
				"name": "Hello World;v1.0.0",
				"domain": {
					"_key": "bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle",
					"name": "GLOBAL.My Project"
				},
				"parent_key": "",
				"predecessor_key": ""
			},
			{
				"_key": "bafkreie77ros2gduaq2mkji5f2deckk2mkgqw4pyveumrwxjzcuzgkda3u",
				"name": "FooBar;v1.0.0",
				"domain": {
					"_key": "bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle",
					"name": "GLOBAL.My Project"
				},
				"parent_key": "",
				"predecessor_key": ""
			}
		]
	}`)

	expected := `{"components":[{"_key":"bafkreieu66waq6jcefgbaxlwkeg6cnqoj5zlc63wghddh3ngtzh7olp37u"},{"_key":"bafkreie77ros2gduaq2mkji5f2deckk2mkgqw4pyveumrwxjzcuzgkda3u"}]}`

	var comps2nft Components // define user object to marshal into

	json.Unmarshal(jsonObj, &comps2nft) // convert json string into the user object
	comps2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", comps2nft.Key, comps2nft.NftJson)
	assert.Equal(t, comps2nft.NftJson, expected, "check nft json against expected results")

	var nft2comps Components // define user object to marshal into

	nft2comps.NftJson = expected     // set the nft json
	nft2comps.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2comps.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, nft2comps.NftJson, expected, "check unmarshalled user against expected results")

}
