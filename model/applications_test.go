package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestApplications(t *testing.T) {

	jsonObj := []byte(`{
		"applications": [{
			"name": "Hello App;v1",
			"domain": {
			  "name": "GLOBAL.My Project"
			},
			"parent_key": "",
			"predecessor_key": "",
			"deployments": [121]
		  }
		]
	}`)

	expected := "{\"applications\":[{\"deployments\":[121],\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"Hello App;v1\"}],\"objtype\":\"Applications\"}"
	expectedCid := "bafkreia37rx5tis4rucgopt4wekpxtquzmitxmzhj7s6aa7th3mayw7xly"

	// define user object to marshal into
	obj := NewApplications()

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
