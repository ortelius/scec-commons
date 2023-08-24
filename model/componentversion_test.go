package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestComponentVersion(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "ComponentVersion",
		"name": "Hello World;v1.0.0",
		"domain": {
			"objtype": "Domain",
			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": ""
	  }`)

	expected := "{\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"Hello World;v1.0.0\",\"objtype\":\"ComponentVersion\"}"
	expectedCid := "bafkreieouznpz5eey2ygrvia7ocaa44n7nsqx63ymvajvireeljuaiolom"

	// define user object to marshal into
	obj := NewComponentVersion()

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
