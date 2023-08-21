package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestComponentVersion(t *testing.T) {

	jsonObj := []byte(`{
		"name": "Hello World;v1.0.0",
		"domain": {
			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": ""
	  }`)

	expected := "{\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"Hello World;v1.0.0\",\"objtype\":\"ComponentVersion\"}"
	expectedCid := "bafkreig4b2yok2qfyfti7x2j574uef4b7yfqrbwzj26rarvo7dexkczbxq"

	// define user object to marshal into
	var obj ComponentVersion

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
