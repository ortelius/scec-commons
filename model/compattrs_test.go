package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestCompAttrs(t *testing.T) {

	jsonObj := []byte(`{

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
	  }`)

	expected := "{\"builddate\":\"0001-01-01T00:00:00Z\",\"gitbranchcreatetimestamp\":\"0001-01-01T00:00:00Z\",\"gitcommittimestamp\":\"0001-01-01T00:00:00Z\",\"objtype\":\"CompAttrs\",\"serviceowner\":{\"domain\":{\"name\":\"\"},\"name\":\"\"}}"
	expectedCid := "bafkreidyx4arxs35uewubsyhqgvuuozcacdjmkocwz3fsyysb7w3gi6jiy"

	// define user object to marshal into
	var obj CompAttrs

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
