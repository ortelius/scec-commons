package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompAttrs(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreibpyqhke4wxmtfjyp2non4gmjyatz5nfhbknvfzfmh2updwtsgd24",
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
		  "_key": "bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi",
		  "name": "admin",
		  "domain": {
			"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
			"name": "GLOBAL"
		  },
		  "email": "admin@ortelius.io",
		  "phone": "505-444-5566",
		  "realname": "Ortelius Admin"
		},
		"slackchannel": "https://myproject.slack.com/444aaa"
	  }`)

	expected := `{"builddate":"Mon Jan 31 16:18:26 2022","buildurl":"https://circleci.com/gh/ortelius/store-cartservice/178","chart":"chart/ms-cartservice","chartnamespace":"default","chartrepo":"msproject/ms-chartservice","chartrepourl":"https://helm.msprogject/stable/msproject/ms-chartservice","chartversion":"1.0.0","discordchannel":"https://discord.gg/A4hx3","dockerrepo":"myproject/ms-chartservice","dockersha":"5d3d677e1","dockertag":"v1.0.0","gitcommit":"2adc111","gitrepo":"msproject/ms-chartservice","gittag":"main","giturl":"https://github.com/msproject/ms-chartservice","objtype":"CompAttr","pagerdutybusinessurl":"https://pagerduty.com/business/ms-chartservice","pagerdutyurl":"https://pagerduty.com/business/ms-chartservice","serviceowner":{"_key":"bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi"},"slackchannel":"https://myproject.slack.com/444aaa"}`

	var compattrs2nft CompAttrs // define user object to marshal into

	json.Unmarshal(jsonObj, &compattrs2nft) // convert json string into the user object
	compattrs2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", compattrs2nft.Key, compattrs2nft.NftJson)
	assert.Equal(t, compattrs2nft.NftJson, expected, "check nft json against expected results")

	var nft2compattrs CompAttrs // define user object to marshal into

	nft2compattrs.NftJson = expected     // set the nft json
	nft2compattrs.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2compattrs.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, nft2compattrs.NftJson, expected, "check unmarshalled user against expected results")

}
