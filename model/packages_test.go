package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestPackages(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "Packages",
		"packages": [{
				"objtype": "Package",
				"purl": "pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
				"name": "libc-bin",
				"version": "2.19.18+deb8u7",
				"license": "GP-2.0",
				"cve": "OSVDEV-1",
				"summary": "test cve"
			},
			{
				"objtype": "Package",
				"purl": "pkg:deb/debian/libcpp-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
				"name": "libcpp-bin",
				"version": "2.19.18+deb8u7",
				"license": "GP-2.0",
				"cve": "OSVDEV-2",
				"summary": "another test cve"
			}
		]
	}`)

	expected := "{\"objtype\":\"Packages\",\"packages\":[{\"cve\":\"OSVDEV-1\",\"license\":\"GP-2.0\",\"name\":\"libc-bin\",\"objtype\":\"Package\",\"purl\":\"pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"summary\":\"test cve\",\"version\":\"2.19.18+deb8u7\"},{\"cve\":\"OSVDEV-2\",\"license\":\"GP-2.0\",\"name\":\"libcpp-bin\",\"objtype\":\"Package\",\"purl\":\"pkg:deb/debian/libcpp-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"summary\":\"another test cve\",\"version\":\"2.19.18+deb8u7\"}]}"
	expectedCid := "bafkreifkosx2sdtbq7qqgpsrvneea2qmbbkujej5tmjmyuw5usfi6un2wa"

	// define user object to marshal into
	obj := NewPackages()

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
