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

	expected := "{\"auditlog\":{\"objtype\":\"AuditLog\"},\"components\":{\"components\":[{\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"Hello World;v1.0.0\"},{\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"FooBar;v1.0.0\"}],\"objtype\":\"Components\"},\"created\":\"2023-04-23T10:20:30.4+02:30\",\"creator\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"Hello App;v1\",\"objtype\":\"ApplicationVersionDetails\",\"owner\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"}}"
	expectedCid := "bafkreigujh46iovijn2r4bdjiu3a5qhxbx3jngytrhnlu5bxgtanalthg4"

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
