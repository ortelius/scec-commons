package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeployment(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreif2i6j3nd5blc6mbs4lfbf3ksjsa5xl2zierqjwmkerrqzcep6ie4",
		"environment": {
			"_key": "bafkreibajlefuhj4rbcmopy6c26riuyhui7kg5znd3ghgqznmuseuhwjea",
			"name": "Development",
			"domain": "GLOBAL.My Project",
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
		},
		"application": {
			"_key": "bafkreia4ioz2a6o3w5ijarqbxwfixcmevlqukjd4bndkw4bj7vosrjqfh4",
			"name": "Hello App;v1",
			"domain": {
				"_key": "bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle",
				"name": "GLOBAL.My Project"
			},
			"parent_key": "",
			"predecessor_key": "",
			"deployments": [121]
		},
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
		],
		"starttime": "2023-04-23T10:20:30.400+02:30",
		"endtime": "2023-04-23T10:30:30.400+02:30",
		"result": 0,
		"deploynum": 100
	}`)

	expected := `{"application":{"_key":"bafkreia4ioz2a6o3w5ijarqbxwfixcmevlqukjd4bndkw4bj7vosrjqfh4"},"components":{"_key":"bafkreicecnx2gvntm6fbcrvnc336qze6st5u7qq7457igegamd3bzkx7ri"},"deploynum":100,"endtime":"2023-04-23T10:30:30.4+02:30","environment":{"_key":"bafkreibgofundww2xdhtdu63wgn3tp76xrccodqq6umkv5xfneifnb4m2a"},"objtype":"Deployment","starttime":"2023-04-23T10:20:30.4+02:30"}`

	var deployment2nft Deployment // define user object to marshal into

	json.Unmarshal(jsonObj, &deployment2nft) // convert json string into the user object
	deployment2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", deployment2nft.Key, deployment2nft.NftJSON)
	assert.Equal(t, expected, deployment2nft.NftJSON, "check nft json against expected results")

	var nft2deployment Deployment // define user object to marshal into

	nft2deployment.NftJSON = expected     // set the nft json
	nft2deployment.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2deployment.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2deployment.NftJSON, "check unmarshalled against expected results")

}
