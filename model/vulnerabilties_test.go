package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestVulnerabilities(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "Vulnerabilities",
		"vulnerabilties": [{
			    "objtype": "Vulnerabilities",
				"name": "CVE-1823"
			},
			{
				"objtype": "Vulnerabilities",
				"name": "CVE-1824"
			}
		]

	}`)

	expected := "{\"objtype\":\"Vulnerabilities\",\"vulnerabilties\":[{\"name\":\"CVE-1824\",\"objtype\":\"Vulnerabilities\"},{\"name\":\"CVE-1823\",\"objtype\":\"Vulnerabilities\"}]}"
	expectedCid := "bafkreihwylfo22ratdjgkstonuhe7e2dfnxzofiv2ooyfdwbyaoqwipilq"

	// define user object to marshal into
	obj := NewVulnerabilities()

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
