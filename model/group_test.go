package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreigqoe7u25fslc5vf54xrjrdw3noxy5cq4bnggsh4pb7i3wdl47ndu",
		"name": "Administrators",
		"domain": {
		  "_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
		  "name": "GLOBAL"
		}
	  }`)

	expected := `{"domain":{"_key":"bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty"},"name":"Administrators","objtype":"Group"}`

	var group2nft Group // define user object to marshal into

	json.Unmarshal(jsonObj, &group2nft) // convert json string into the user object
	group2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", group2nft.Key, group2nft.NftJSON)
	assert.Equal(t, expected, group2nft.NftJSON, "check nft json against expected results")

	var nft2user Group // define user object to marshal into

	nft2user.NftJSON = expected     // set the nft json
	nft2user.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2user.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2user.NftJSON, "check unmarshalled user against expected results")

}
