package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestDomain(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "Domain",
		"name": "GLOBAL"
	  }`)

	expected := "{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"}"
	expectedCid := "bafkreihfbbgm73jj22esx7wni4npeqcwr2tn6zobgfwkjkqnua4em5fgja"

	// define user object to marshal into
	obj := NewDomain()

	// convert json string into the user object
	json.Unmarshal(jsonObj, obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(obj)
	// fmt.Println(cid)
	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")
}
