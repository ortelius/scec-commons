package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestGroupsForUser(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "GroupsForUser",
		"user": "admin",
		"groups": [ "users", "administrators" ]
	  }`)

	expected := "{\"groups\":[\"administrators\",\"users\"],\"objtype\":\"GroupsForUser\",\"user\":\"admin\"}"
	expectedCid := "bafkreihks5vdn3ebfephucdre2t4lpus6v75lhsplutatvtqjmq6soi5my"

	// define user object to marshal into
	obj := NewGroupsForUser()

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
