package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestApplicationVersionDetails(t *testing.T) {

	jsonObj := []byte(`{

		"name": "Hello App;v1",
		"domain": {

			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": "",
		"deployments": [],
		"owner": {

			"name": "admin",
			"domain": {

				"name": "GLOBAL"
			},
			"email": "admin@ortelius.io",
			"phone": "505-444-5566",
			"realname": "Ortelius Admin"
		},
		"creator": {

			"name": "admin",
			"domain": {

				"name": "GLOBAL"
			},
			"email": "admin@ortelius.io",
			"phone": "505-444-5566",
			"realname": "Ortelius Admin"
		},
		"created": "2023-04-23T10:20:30.400+02:30",
		"components": {

				"components": [{

						"name": "Hello World;v1.0.0",
						"domain": {

							"name": "GLOBAL.My Project"
						},
						"parent_key": "",
						"predecessor_key": ""
					},
					{

						"name": "FooBar;v1.0.0",
						"domain": {

							"name": "GLOBAL.My Project"
						},
						"parent_key": "",
						"predecessor_key": ""
					}
				]
		},
		"auditlog": {

			"auditlog": []
		}
	}`)

	expected := "{\"components\":{\"components\":[{\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"Hello World;v1.0.0\"},{\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"FooBar;v1.0.0\"}]},\"created\":\"2023-04-23T10:20:30.4+02:30\",\"creator\":{\"domain\":{\"name\":\"GLOBAL\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"Hello App;v1\",\"objtype\":\"ApplicationVersionDetails\",\"owner\":{\"domain\":{\"name\":\"GLOBAL\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"}}"
	expectedCid := "bafkreib2rmpefrysoitdg7mprkfsol5jbgo6ejdtzxurbl5iub7gqg7mea"

	// define user object to marshal into
	var obj ApplicationVersionDetails

	// convert json string into the user object
	json.Unmarshal(jsonObj, &obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(obj)
	// 	fmt.Println(cid)
	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")

}
