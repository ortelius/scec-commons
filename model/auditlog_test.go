package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestAuditLog(t *testing.T) {

	jsonObj := []byte(`{

		"auditlog": [{
				"objtype": "AuditLog",
				"action": "Created",
				"user": {
					"objtype": "User",
					"domain": {
						"objtype": "Domain",
						"name": "GLOBAL"
					},
					"email": "admin@ortelius.io",
					"name": "admin",
					"phone": "505-444-5566",
					"realname": "Ortelius Admin"
				},
				"when": "2023-04-23T10:20:30.400+02:30"
			},
			{
				"objtype": "AuditLog",
				"action": "Updated",
				"user": {
					"objtype": "User",
					"domain": {
						"objtype": "Domain",
						"name": "GLOBAL"
					},
					"email": "admin@ortelius.io",
					"name": "admin",
					"phone": "505-444-5566",
					"realname": "Ortelius Admin"
				},
				"when": "2023-05-23T10:20:30.400+02:30"
			}
		]
	}`)

	expected := "{\"auditlog\":[{\"User\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"action\":\"Created\",\"objtype\":\"AuditLog\",\"when\":\"2023-04-23T10:20:30.4+02:30\"},{\"User\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"action\":\"Updated\",\"objtype\":\"AuditLog\",\"when\":\"2023-05-23T10:20:30.4+02:30\"}],\"objtype\":\"AuditLog\"}"
	expectedCid := "bafkreiep66t3sp77nricy2vq4ms3eyxbhyp3vreqzgp4g4rtx2v5v3pony"

	// define user object to marshal into
	obj := NewAuditLog()

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
