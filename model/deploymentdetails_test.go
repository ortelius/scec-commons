package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestDeploymentDetails(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "DeploymentDetails",
		"log": [
			"Starting",
			"Finished"
		],
		"deployment": {
			"objtype": "Deployment",
			"environment": {
				"objtype": "Environment",
				"name": "Development",
				"domain": "GLOBAL.My Project",
				"owner": {
					"objtype": "User",
					"name": "admin",
					"domain": {
						"objtype": "Domain",
						"name": "GLOBAL"
					},
					"email": "admin@ortelius.io",
					"phone": "505-444-5566",
					"realname": "Ortelius Admin"
				},
				"creator": {
					"objtype": "User",
					"name": "admin",
					"domain": {
						"objtype": "Domain",
						"name": "GLOBAL"
					},
					"email": "admin@ortelius.io",
					"phone": "505-444-5566",
					"realname": "Ortelius Admin"
				},
				"created": "2023-04-23T10:20:30.400+02:30"
			},
			"application": {
				"objtype": "Application",
				"name": "Hello App;v1",
				"domain": {
					"objtype": "Domain",
					"name": "GLOBAL.My Project"
				},
				"parent_key": "",
				"predecessor_key": "",
				"deployments": [121]
			},
			"components": [{
					"objtype": "Component",
					"name": "Hello World;v1.0.0",
					"domain": {
						"objtype": "Domain",
						"name": "GLOBAL.My Project"
					},
					"parent_key": "",
					"predecessor_key": ""
				},
				{
					"objtype": "Component",
					"name": "FooBar;v1.0.0",
					"domain": {
						"objtype": "Domain",
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

	expected := "{\"deployment\":{\"application\":{\"deployments\":[121],\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"Hello App;v1\",\"objtype\":\"Application\"},\"components\":{\"objtype\":\"Components\"},\"deploynum\": 100,\"endtime\":\"2023-04-23T10:30:30.4+02:30\",\"environment\":{\"created\":\"2023-04-23T10:20:30.4+02:30\",\"creator\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"domain\":{\"name\":\"\",\"objtype\":\"Domain\"},\"name\":\"Development\",\"objtype\":\"Environment\",\"owner\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"}},\"objtype\":\"Deployment\",\"starttime\":\"2023-04-23T10:20:30.4+02:30\"},\"log\":[\"Finished\",\"Starting\"],\"objtype\":\"DeploymentDetails\"}"
	expectedCid := "bafkreihriwzmphzmhn3sgda3yh5ramcrdjt4hropshohdbueyvva3k5lv4"

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
