package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "Group",
		"name": "Administrators",
		"domain": {
		  "objtype": "Domain",
		  "name": "GLOBAL"
		}
	  }`)

	expected := "{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"name\":\"Administrators\",\"objtype\":\"Group\"}"
	expectedCid := "bafkreiavyiihdr4k5ijrzkuva2byt2kqpnpxz5rwovsobag7tcwx5gsknu"

	// define user object to marshal into
	obj := NewGroup()

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
