package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestDeploymentDetails(t *testing.T) {

	jsonObj := []byte(`{

		"log": [
			"Starting",
			"Finished"
		],
		"deployment": {

			"environment": {

				"name": "Development",
				"domain": "GLOBAL.My Project",
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
				"created": "2023-04-23T10:20:30.400+02:30"
			},
			"application": {

				"name": "Hello App;v1",
				"domain": {

					"name": "GLOBAL.My Project"
				},
				"parent_key": "",
				"predecessor_key": "",
				"deployments": [121]
			},
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
			],
			"starttime": "2023-04-23T10:20:30.400+02:30",
			"endtime": "2023-04-23T10:30:30.400+02:30",
			"result": 0,
			"deploynum": 100
		}
	}`)

	expected := "{\"deployment\":{\"application\":{\"deployments\":[121],\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"Hello App;v1\"},\"deploynum\": 100,\"endtime\":\"2023-04-23T10:30:30.4+02:30\",\"environment\":{\"created\":\"2023-04-23T10:20:30.4+02:30\",\"creator\":{\"domain\":{\"name\":\"GLOBAL\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"domain\":{\"name\":\"\"},\"name\":\"Development\",\"owner\":{\"domain\":{\"name\":\"GLOBAL\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"}},\"starttime\":\"2023-04-23T10:20:30.4+02:30\"},\"log\":[\"Finished\",\"Starting\"],\"objtype\":\"DeploymentDetails\"}"
	expectedCid := "bafkreiaybc24p6dgqx5zgaowgimx7vbckkmjodxjviexmnwglucqhirvqa"

	// define user object to marshal into
	var obj DeploymentDetails

	// convert json string into the user object
	json.Unmarshal(jsonObj, &obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(&obj)
	// 	fmt.Println(cid)
	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")

}
