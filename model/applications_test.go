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
			"objtype": "Applications",
			"name": "Hello App;v1",
			"domain": {
			  "objtype": "Domain",
			  "name": "GLOBAL.My Project"
			},
			"parent_key": "",
			"predecessor_key": "",
			"deployments": [121]
		  }
		]
	}`)

	expected := "{\"applications\":[{\"deployments\":[121],\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"Hello App;v1\",\"objtype\":\"Applications\"}],\"objtype\":\"Applications\"}"
	expectedCid := "bafkreiamsg6kb47toamrydimhno2orwireuqpjeb6uvn5ssnqch4feaol4"

	// define user object to marshal into
	obj := NewApplications()

	// convert json string into the user object
	json.Unmarshal(jsonObj, obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(obj)
	// fmt.Println(dbStr)
	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")
}
