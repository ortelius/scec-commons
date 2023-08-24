package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestApplicationVersion(t *testing.T) {

	jsonObj := []byte(`{
		"name": "Hello App;v1",
		"objtype": "ApplicationVersion",
		"domain": {
		  "objtype": "Domain",
		  "name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": "",
		"deployments": [121]
	  }`)

	expected := "{\"deployments\":[121],\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"Hello App;v1\",\"objtype\":\"ApplicationVersion\"}"
	expectedCid := "bafkreidpcugtjirjddxbkove3lzac4kaaxzlftpet54546sjdzvefwjyde"

	// define user object to marshal into
	obj := NewApplicationVersion()

	// convert json string into the user object
	json.Unmarshal(jsonObj, obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(obj)

	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")

}
