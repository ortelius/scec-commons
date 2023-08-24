package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestApplicationVersionDetails(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "ApplicationVersionDetails",
		"name": "Hello App;v1",
		"domain": {
			"objtype": "Domain",
			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": "",
		"deployments": [],
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
		"created": "2023-04-23T10:20:30.400+02:30",
		"components": {
				"objtype": "Components",
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
				]
		},
		"auditlog": {
			"objtype": "AuditLog",
			"auditlog": []
		}
	}`)

	expected := "{\"auditlog\":{\"objtype\":\"AuditLog\"},\"components\":{\"components\":[{\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"Hello World;v1.0.0\",\"objtype\":\"Component\"},{\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"FooBar;v1.0.0\",\"objtype\":\"Component\"}],\"objtype\":\"Components\"},\"created\":\"2023-04-23T10:20:30.4+02:30\",\"creator\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"Hello App;v1\",\"objtype\":\"ApplicationVersionDetails\",\"owner\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"}}"
	expectedCid := "bafkreicwgwyeamaaioszvmizco7z3cnwv4g23bftomqvx5ntazgvrkvhr4"

	// define user object to marshal into
	obj := NewApplicationVersionDetails()

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
