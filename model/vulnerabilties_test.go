package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestVulnerabilities(t *testing.T) {

	jsonObj := []byte(`{

		"vulnerabilties": [{

				"name": "CVE-1823"
			},
			{

				"name": "CVE-1824"
			}
		]

	}`)

	expected := "{\"objtype\":\"Vulnerabilities\",\"vulnerabilties\":[{\"name\":\"CVE-1823\"},{\"name\":\"CVE-1824\"}]}"
	expectedCid := "bafkreihf2jx4bieq6slnzrfrg4xdmnntbdmqtbyefviuy2apt6jhzmc5w4"

	// define user object to marshal into
	var obj Vulnerabilities

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
