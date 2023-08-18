package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestComponentVersionDetails(t *testing.T) {

	jsonObj := []byte(`{

		"name": "Hello World;v1.0.0",
		"domain": {

			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": "",
		"owner": {

			"name": "admin",
			"domain": {

				"name": "GLOBAL"
			},
			"email": "admin@ortelius.io",
			"phone": "505-444-5566",
			"realname": "Ortelius Admin"
		},
		"creator": {

			"name": "admin",
			"domain": {

				"name": "GLOBAL"
			},
			"email": "admin@ortelius.io",
			"phone": "505-444-5566",
			"realname": "Ortelius Admin"
		},
		"created": "2023-04-23T10:20:30.400+02:30",
		"comptype": "docker",
		"packages": {

			"packages": [{

					"purl": "pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
					"name": "libc-bin",
					"version": "2.19.18+deb8u7",
					"license_key": 23,
					"license": "GP-2.0"
				},
				{

					"purl": "pkg:deb/debian/libcpp-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
					"name": "libcpp-bin",
					"version": "2.19.18+deb8u7",
					"license_key": 23,
					"license": "GP-2.0"
				}
			]
		},
		"vulnerabilties": {

			"vulnerabilties": [{

					"name": "CVE-1823"
				},
				{

					"name": "CVE-1824"
				}
			]
		},
		"readme": {

			"content": ["# README", "## Sample"]
		},
		"license": {

			"content": ["# Apache 2", "## Summary"]
		},
		"swagger": {

			"content": ["# Rest APIs", "## GET /user"]
		},
		"applications": {

			"applications": []
		},
		"providing": {

			"provides": ["/user"]
		},
		"consuming": {

			"consumes": ["/user"]
		},
		"attrs": {

			"builddate": "Mon Jan 31 16:18:26 2022",
			"build_key": "178",
			"buildurl": "https://circleci.com/gh/ortelius/store-cartservice/178",
			"chart": "chart/ms-cartservice",
			"chartnamespace": "default",
			"chartrepo": "msproject/ms-chartservice",
			"chartrepourl": "https://helm.msprogject/stable/msproject/ms-chartservice",
			"chartversion": "1.0.0",
			"discordchannel": "https://discord.gg/A4hx3",
			"dockerrepo": "myproject/ms-chartservice",
			"dockersha": "5d3d677e1",
			"dockertag": "v1.0.0",
			"gitcommit": "2adc111",
			"gitrepo": "msproject/ms-chartservice",
			"gittag": "main",
			"giturl": "https://github.com/msproject/ms-chartservice",
			"hipchatchannel": "",
			"pagerdutybusinessurl": "https://pagerduty.com/business/ms-chartservice",
			"pagerdutyurl": "https://pagerduty.com/business/ms-chartservice",
			"serviceowner": {

				"name": "admin",
				"domain": {

					"name": "GLOBAL"
				},
				"email": "admin@ortelius.io",
				"phone": "505-444-5566",
				"realname": "Ortelius Admin"
			},
			"slackchannel": "https://myproject.slack.com/444aaa"
		}
	}`)

	expected := "{\"attrs\":{\"builddate\":\"0001-01-01T00:00:00Z\",\"gitbranchcreatetimestamp\":\"0001-01-01T00:00:00Z\",\"gitcommittimestamp\":\"0001-01-01T00:00:00Z\",\"serviceowner\":{\"domain\":{\"name\":\"\"},\"name\":\"\"}},\"comptype\":\"docker\",\"consuming\":{\"consumes\":[\"/user\"]},\"created\":\"2023-04-23T10:20:30.4+02:30\",\"creator\":{\"domain\":{\"name\":\"GLOBAL\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"domain\":{\"name\":\"GLOBAL.My Project\"},\"license\":{\"content\":[\"# Apache 2\",\"## Summary\"]},\"name\":\"Hello World;v1.0.0\",\"objtype\":\"ComponentVersionDetails\",\"owner\":{\"domain\":{\"name\":\"GLOBAL\"},\"email\":\"admin@ortelius.io\",\"name\":\"admin\",\"phone\":\"505-444-5566\",\"realname\":\"Ortelius Admin\"},\"packages\":{\"packages\":[{\"license\":\"GP-2.0\",\"name\":\"libc-bin\",\"purl\":\"pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"version\":\"2.19.18+deb8u7\"},{\"license\":\"GP-2.0\",\"name\":\"libcpp-bin\",\"purl\":\"pkg:deb/debian/libcpp-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"version\":\"2.19.18+deb8u7\"}]},\"providing\":{\"provides\":[\"/user\"]},\"readme\":{\"content\":[\"# README\",\"## Sample\"]},\"swagger\":{\"content\":[\"# Rest APIs\",\"## GET /user\"]},\"vulnerabilties\":{\"vulnerabilties\":[{\"name\":\"CVE-1823\"},{\"name\":\"CVE-1824\"}]}}"
	expectedCid := "bafkreicmf3egdhty2kiimwh42qqgumuqnqgi3xvt7ziho2n2t5tnremb4q"

	// define user object to marshal into
	var obj ComponentVersionDetails

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
