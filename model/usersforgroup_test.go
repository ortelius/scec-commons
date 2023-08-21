package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestUsersForGroup(t *testing.T) {

	jsonObj := []byte(`{

		"group": "administrators",
		"users": [ "testadmin", "admin" ]
	  }`)

	expected := "{\"group\":\"administrators\",\"objtype\":\"UsersForGroup\",\"users\":[\"admin\",\"testadmin\"]}"
	expectedCid := "bafkreigkc6rwtkltn2dhqivs2mfvlzk6djkajfnogffnmubzlqjmioea7m"

	// define user object to marshal into
	var obj UsersForGroup

	// convert json string into the user object
	json.Unmarshal(jsonObj, &obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(&obj)
	// 	fmt.Println(cid)
	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")

}
