package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {

	jsonObj := []byte(`{

		"name": "Administrators",
		"domain": {

		  "name": "GLOBAL"
		}
	  }`)

	expected := "{\"domain\":{\"name\":\"GLOBAL\"},\"name\":\"Administrators\",\"objtype\":\"Group\"}"
	expectedCid := "bafkreiam46rnoswipyv7qyysnow5rxuaoq2wsatd2fwynciwkm3sg5ocra"

	// define user object to marshal into
	var obj Group

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
