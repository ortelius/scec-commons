package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "Package",
		"purl": "pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
		"name": "libc-bin",
		"version": "2.19.18+deb8u7",
		"license": "GP-2.0",
		"cve": "OSVDEV-1",
		"summary": "test cve"
	}`)

	expected := "{\"cve\":\"OSVDEV-1\",\"license\":\"GP-2.0\",\"name\":\"libc-bin\",\"objtype\":\"Package\",\"purl\":\"pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"summary\":\"test cve\",\"version\":\"2.19.18+deb8u7\"}"
	expectedCid := "bafkreialuff3qicjkv3gydigmea7k3bz4facigzk7fqppmii6diqokjida"

	// define user object to marshal into
	obj := NewPackage()

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
