package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestAuditRecord(t *testing.T) {

	jsonObj := []byte(`{

		"action": "Created",
		"user": {

		  "domain": {

			"name": "GLOBAL"
		  },
		  "email": "admin@ortelius.io",
		  "name": "admin",
		  "phone": "505-444-5566",
		  "realname": "Ortelius Admin"
		},
		"when": "2023-04-23T10:20:30.400+02:30"
	  }`)

	expected := "{\"User\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"action\":\"Created\",\"objtype\":\"AuditRecord\",\"when\":\"2023-04-23T10:20:30.4+02:30\"}"
	expectedCid := "bafkreicq4wlsybz4qpsdxqugue4em7z7h5oolgecbecmoao75pdhtrhekq"

	// define user object to marshal into
	obj := NewAuditRecord()

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
