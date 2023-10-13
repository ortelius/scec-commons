package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestComponentVersionDetails(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "ComponentVersionDetails",
		"name": "Hello World;v1.0.0",
		"domain": {
			"objtype": "Domain",
			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": "",
		"owner": {
			"objtype": "User",
			"name": "admin",
			"domain": {
				"objtype": "Domain",
				"name": "GLOBAL"
			},
			"email": "admin@ortelius.io",
			"phone": "505-444-5566",
			"realname": "Ortelius Admin"
		},
		"creator": {
			"objtype": "User",
			"name": "admin",
			"domain": {
				"objtype": "Domain",
				"name": "GLOBAL"
			},
			"email": "admin@ortelius.io",
			"phone": "505-444-5566",
			"realname": "Ortelius Admin"
		},
		"created": "2023-04-23T10:20:30.400+02:30",
		"comptype": "docker",
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
		}],
		"readme": {
			"objtype": "Readme",
			"content": ["# README", "## Sample"]
		},
		"license": {
			"objtype": "License",
			"content": ["# Apache 2", "## Summary"]
		},
		"swagger": {
			"objtype": "Swagger",
			"content": ["# Rest APIs", "## GET /user"]
		},
		"applications": {
			"objtype": "Applications",
			"applications": []
		},
		"providing": {
			"objtype": "Providing",
			"provides": ["/user"]
		},
		"consuming": {
			"objtype": "Consuming",
			"consumes": ["/user"]
		}
	}`)

	expected := "{\"attrs\":{\"builddate\":\"0001-01-01T00:00:00Z\",\"gitbranchcreatetimestamp\":\"0001-01-01T00:00:00Z\",\"gitcommittimestamp\":\"0001-01-01T00:00:00Z\",\"objtype\":\"CompAttrs\",\"serviceowner\":{\"domain\":{\"name\":\"\",\"objtype\":\"Domain\"},\"name\":\"\",\"objtype\":\"User\"}},\"autditlog\":{\"objtype\":\"AuditLog\"},\"comptype\":\"docker\",\"consuming\":{\"consumes\":[\"/user\"],\"objtype\":\"Consuming\"},\"created\":\"2023-04-23T10:20:30.4+02:30\",\"creator\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"domain\":{\"name\":\"GLOBAL.My Project\",\"objtype\":\"Domain\"},\"license\":{\"content\":[\"# Apache 2\",\"## Summary\"],\"objtype\":\"License\"},\"name\":\"Hello World;v1.0.0\",\"objtype\":\"ComponentVersionDetails\",\"owner\":{\"domain\":{\"name\":\"GLOBAL\",\"objtype\":\"Domain\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"objtype\":\"User\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"packages\":[{\"cve\":\"OSVDEV-1\",\"license\":\"GP-2.0\",\"name\":\"libc-bin\",\"objtype\":\"Package\",\"purl\":\"pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"summary\":\"test cve\",\"version\":\"2.19.18+deb8u7\"},{\"cve\":\"OSVDEV-2\",\"license\":\"GP-2.0\",\"name\":\"libcpp-bin\",\"objtype\":\"Package\",\"purl\":\"pkg:deb/debian/libcpp-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"summary\":\"another test cve\",\"version\":\"2.19.18+deb8u7\"}],\"providing\":{\"objtype\":\"Providing\",\"provides\":[\"/user\"]},\"readme\":{\"content\":[\"# README\",\"## Sample\"],\"objtype\":\"Readme\"},\"swagger\":{\"content\":[\"# Rest APIs\",\"## GET /user\"],\"objtype\":\"Swagger\"}}"
	expectedCid := "bafkreia25sipepelyaowr3oycgvtd4rqc4xun67p5zrazng2c7yrf2pyt4"

	// define user object to marshal into
	obj := NewComponentVersionDetails()

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
