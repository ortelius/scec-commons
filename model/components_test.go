package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestComponents(t *testing.T) {

	jsonObj := []byte(`{

		"components": [{

				"name": "Hello World;v1.0.0",
				"domain": {

					"name": "GLOBAL.My Project"
				},
				"parent_key": "",
				"predecessor_key": ""
			},
			{

				"name": "FooBar;v1.0.0",
				"domain": {

					"name": "GLOBAL.My Project"
				},
				"parent_key": "",
				"predecessor_key": ""
			}
		]
	}`)

	expected := "{\"components\":[{\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"Hello World;v1.0.0\"},{\"domain\":{\"name\":\"GLOBAL.My Project\"},\"name\":\"FooBar;v1.0.0\"}],\"objtype\":\"Components\"}"
	expectedCid := "bafkreif6vuydubg7cam3imzikxw5lh6l6uu7s5hxkvbzda2baculds7m6q"

	// define user object to marshal into
	obj := NewComponents()

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
