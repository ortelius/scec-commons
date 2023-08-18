package model

import (
	"encoding/json"
	"testing"

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

	expected := `{"domain":{"name":"GLOBAL.My Project"},"name":"Hello World;v1.0.0"}`
	expectedCid := "bafkreidbz3f4apbvla5whzprg7lofwiqefbeh6a44mwgmkpsezykzlpnmu"

	// define user object to marshal into
	var obj ComponentVersion

	// convert json string into the user object
	json.Unmarshal(jsonObj, &obj)

	// create all cids for the json string
	cid, _ := MakeNFT(obj)
	assert.Equal(t, cid, expectedCid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := MakeJSON(cid)
	assert.Equal(t, jsonStr, expected, "check persisted cid json with test json string")
}
