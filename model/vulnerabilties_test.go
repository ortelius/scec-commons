package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVulnerabilities(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreigt6hy7df5mqqo7akcvkmu62aovlrv7sm7mqw6udp2f3777czh5i4",
		"vulnerabilties": [{
				"_key": "bafkreie4t6xqy65fsykdtbz5rcwjivii7rpapjq55zwnpnbl6uj6dysdbu",
				"name": "CVE-1823"
			},
			{
				"_key": "bafkreifpc2htaxkfttbyzywe4nsn2kbicq6w5rj2sqf2lr6nw5qw5qqzka",
				"name": "CVE-1824"
			}
		]

	}`)

	expected := `{"vulnerabilties":[{"_key":"bafkreie4t6xqy65fsykdtbz5rcwjivii7rpapjq55zwnpnbl6uj6dysdbu"},{"_key":"bafkreifpc2htaxkfttbyzywe4nsn2kbicq6w5rj2sqf2lr6nw5qw5qqzka"}]}`

	var vulns2nft Vulnerabilities // define user object to marshal into

	json.Unmarshal(jsonObj, &vulns2nft)       // convert json string into the user object
	nftJSON := vulns2nft.MarshalNFT(cid2json) // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", vulns2nft.Key, vulns2nft.NftJSON)
	assert.Equal(t, expected, nftJSON, "check nft json against expected results")

	var nft2vulns Vulnerabilities // define user object to marshal into

	nft2vulns.Key = vulns2nft.Key           // set the nft json
	nft2vulns.UnmarshalNFT(cid2json)        // convert the json string into the user object
	check := nft2vulns.MarshalNFT(cid2json) // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, check, "check unmarshalled user against expected results")

}
