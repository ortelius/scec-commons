package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {

	jsonObj := []byte(`{

		"purl": "pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
		"name": "libc-bin",
		"version": "2.19.18+deb8u7",
		"license_key": 23,
		"license": "GP-2.0"
	  }`)

	expected := "{\"license\":\"GP-2.0\",\"name\":\"libc-bin\",\"objtype\":\"Package\",\"purl\":\"pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"version\":\"2.19.18+deb8u7\"}"
	expectedCid := "bafkreifwfll3ytl3ibzlvfqn6jqhdubbls5rcnqdh6cwbczgw2vk7xw6gi"

	// define user object to marshal into
	var obj Package

	// convert json string into the user object
	json.Unmarshal(jsonObj, &obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(obj)
	// 	fmt.Println(cid)
	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")

}
