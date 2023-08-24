package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestProviding(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "Providing",
		"provides": ["/user"]
	}`)

	expected := "{\"objtype\":\"Providing\",\"provides\":[\"/user\"]}"
	expectedCid := "bafkreih23ub6vya6ymzu3vrfd5lre24zgnafquqvfxsabzcesrsbwohgzq"

	// define user object to marshal into
	obj := NewProviding()

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
