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

	expected := "{\"deployment\":{\"application\":{\"deployments\":[121],\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"Hello App;v1\",\"objtype\":\"ApplicationVersion\"},\"components\":{\"objtype\":\"Components\"},\"deploynum\": 100,\"endtime\":\"2023-04-23T10:30:30.4+02:30\",\"environment\":{\"created\":\"2023-04-23T10:20:30.4+02:30\",\"creator\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"domain\":{\"name\":\"\",\"objtype\":\"Domain\"},\"name\":\"Development\",\"objtype\":\"Environment\",\"owner\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"}},\"objtype\":\"Deployment\",\"starttime\":\"2023-04-23T10:20:30.4+02:30\"},\"log\":[\"Finished\",\"Starting\"],\"objtype\":\"DeploymentDetails\"}"
	expectedCid := "bafkreibylw7hfkil4v5zyz7aeka46nx7sh7rybp3mmiy3pix455fzkakou"

	// define user object to marshal into
	obj := NewDeploymentDetails()

	// convert json string into the user object
	json.Unmarshal(jsonObj, obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(obj)
	// 	fmt.Println(cid)
	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")

}
