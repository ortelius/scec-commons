package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestLicense(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "License",
		"content": ["# Apache 2", "## Summary"]
	  }`)

	expected := "{\"content\":[\"# Apache 2\",\"## Summary\"],\"objtype\":\"License\"}"
	expectedCid := "bafkreiddm2abcmm7l2xszf7te57no5gmwda344ghq3krwqaapcn6wcxeoi"

	// define user object to marshal into
	obj := NewLicense()

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
