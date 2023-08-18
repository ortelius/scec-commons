package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {

	jsonObj := []byte(`{

		"name": "Development",
		"domain": {

			"name": "GLOBAL.My Project"
		},
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
	  }`)

	expected := "{\"created\":\"2023-04-23T10:20:30.4+02:30\",\"creator\":{\"domain\":{\"name\":\"GLOBAL\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"Development\",\"objtype\":\"Environment\",\"owner\":{\"domain\":{\"name\":\"GLOBAL\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"}}"
	expectedCid := "bafkreiakkfewm3yrjpuyfjungz2uqzy67qptpvt25uw2griuil7qzpswe4"

	// define user object to marshal into
	var obj Environment

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
