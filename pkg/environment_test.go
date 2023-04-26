package pkg

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreia35xvytxqxwbhov5osdiobkladnlkoh4gqq47jfr4gohrre562ta",
		"name": "Development",
		"domain": {
			"_key": "bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle",
			"name": "GLOBAL.My Project"
		},
		"owner": {
		  "_key": "bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi",
		  "name": "admin",
		  "domain": {
			"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
			"name": "GLOBAL"
		  },
		  "email": "admin@ortelius.io",
		  "phone": "505-444-5566",
		  "realname": "Ortelius Admin"
		},
		"creator": {
		  "_key": "bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi",
		  "name": "admin",
		  "domain": {
			"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
			"name": "GLOBAL"
		  },
		  "email": "admin@ortelius.io",
		  "phone": "505-444-5566",
		  "realname": "Ortelius Admin"
		},
		"created": "2023-04-23T10:20:30.400+02:30"
	  }`)

	expected := `{"created":"2023-04-23T10:20:30.4+02:30","creator":{"_key":"bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi"},"domain":{"_key":"bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle"},"name":"Development","objtype":"Environment","owner":{"_key":"bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi"}}`

	var env2nft Environment // define user object to marshal into

	json.Unmarshal(jsonObj, &env2nft) // convert json string into the user object
	env2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", env2nft.Key, env2nft.NftJSON)
	assert.Equal(t, expected, env2nft.NftJSON, "check nft json against expected results")

	var nft2env Environment // define user object to marshal into

	nft2env.NftJSON = expected     // set the nft json
	nft2env.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2env.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2env.NftJSON, "check unmarshalled user against expected results")

}
