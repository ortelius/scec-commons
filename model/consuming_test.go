package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestConsuming(t *testing.T) {

	jsonObj := []byte(`{

		"consumes": ["/user"]
	  }`)

	expected := "{\"consumes\":[\"/user\"],\"objtype\":\"Consuming\"}"
	expectedCid := "bafkreiax3vxopxpaz24iqrqnm6f2n4jckz2mstjwt2nc2heursfu56ufoy"

	// define user object to marshal into
	obj := NewConsuming()

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
