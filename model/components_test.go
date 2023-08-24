package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestComponents(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "Components",
		"components": [{
				"objtype": "Component",
				"name": "Hello World;v1.0.0",
				"domain": {
					"objtype": "Domain",
					"name": "GLOBAL.My Project"
				},
				"parent_key": "",
				"predecessor_key": ""
			},
			{
				"objtype": "Component",
				"name": "FooBar;v1.0.0",
				"domain": {
					"objtype": "Domain",
					"name": "GLOBAL.My Project"
				},
				"parent_key": "",
				"predecessor_key": ""
			}
		]
	}`)

	expected := "{\"components\":[{\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"Hello World;v1.0.0\",\"objtype\":\"Component\"},{\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"name\":\"FooBar;v1.0.0\",\"objtype\":\"Component\"}],\"objtype\":\"Components\"}"
	expectedCid := "bafkreifkd3fhlybdw66ifzfoi2u3u4niafsu44ipy2wf4ve5wu3qya7fmy"

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
