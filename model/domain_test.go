package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomain(t *testing.T) {

	jsonObj := []byte(`{
		"name": "GLOBAL"
	  }`)

	expected := `{"name":"GLOBAL"}`
	expectedCid := "bafkreic4j7wtiphcrzz4hpwyqszozkvyfgwfvmmvpblxekjff2nx6f4ism"

	// define user object to marshal into
	var obj Domain

	// convert json string into the user object
	json.Unmarshal(jsonObj, &obj)

	// create all cids for the json string
	cid, _ := MakeNFT(obj)
	// fmt.Println(cid)
	assert.Equal(t, cid, expectedCid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := MakeJSON(cid)
	assert.Equal(t, jsonStr, expected, "check persisted cid json with test json string")
}
