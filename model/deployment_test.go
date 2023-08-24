package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestDeployment(t *testing.T) {

	jsonObj := []byte(`{
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
					"objtype": "Domain,
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
	}`)

	expected := "{\"application\":{\"domain\":{\"name\":\"\",\"objtype\":\"Domain\"},\"name\":\"\",\"objtype\":\"ApplicationVersion\"},\"components\":{\"objtype\":\"Components\"},\"deploynum\": 0,\"endtime\":\"0001-01-01T00:00:00Z\",\"environment\":{\"created\":\"0001-01-01T00:00:00Z\",\"creator\":{\"domain\":{\"name\":\"\",\"objtype\":\"Domain\"},\"name\":\"\",\"objtype\":\"User\"},\"domain\":{\"name\":\"\",\"objtype\":\"Domain\"},\"name\":\"\",\"objtype\":\"Environment\",\"owner\":{\"domain\":{\"name\":\"\",\"objtype\":\"Domain\"},\"name\":\"\",\"objtype\":\"User\"}},\"objtype\":\"Deployment\",\"starttime\":\"0001-01-01T00:00:00Z\"}"
	expectedCid := "bafkreibozc4soiebb5ku7njeza3grrj5v22kxlgnez5ozsvnrrdch3m67u"

	// define user object to marshal into
	obj := NewDeployment()

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
