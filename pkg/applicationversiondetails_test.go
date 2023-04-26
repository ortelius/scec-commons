package pkg

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationVersionDetails(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreibmf6snxkzqqosl4fx4d67unxhq65ekybmsnxehhiwb4cxdnbeduu",
		"name": "Hello App;v1",
		"domain": {
			"_key": "bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle",
			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": "",
		"deployments": [],
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
		"created": "2023-04-23T10:20:30.400+02:30",
		"components": {
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
		},
		"auditlog": {
			"_key": "bafkreicecnx2gvntm6fbcrvnc336qze6st5u7qq7457igegamd3bzkx7ri",
			"auditlog": []
		}
	}`)

	expected := `{"auditlog":{"_key":"bafkreicecnx2gvntm6fbcrvnc336qze6st5u7qq7457igegamd3bzkx7ri"},"components":{"_key":"bafkreibpknufg2ciqkiqlupkglpmlrck5askj3sscieiedwwtwrcffdzwe"},"created":"2023-04-23T10:20:30.4+02:30","creator":{"_key":"bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi"},"domain":{"_key":"bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle"},"name":"Hello App;v1","objtype":"ApplicationVersionDetails","owner":{"_key":"bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi"}}`

	var appver2nft ApplicationVersionDetails // define user object to marshal into

	json.Unmarshal(jsonObj, &appver2nft) // convert json string into the user object
	// fmt.Printf("%+v\n\n", appver2nft)
	appver2nft.MarshalNFT(cid2json) // generate the cid and nft json for user object
	assert.Equal(t, appver2nft.NftJSON, expected, "check nft json against expected results")

	var nft2appver ApplicationVersionDetails // define user object to marshal into

	nft2appver.NftJSON = expected     // set the nft json
	nft2appver.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2appver.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, nft2appver.NftJSON, expected, "check unmarshalled user against expected results")

}
