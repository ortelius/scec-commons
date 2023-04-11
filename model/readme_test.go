package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadme(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreigp3a4djvcp5uidon4366bnggouiw4ghts6ngwfgt5jncbtj6lshy",
		"content": ["# README", "## Sample"]
	  }`)

	expected := `{"content":["# README","## Sample"],"objtype":"Readme"}`

	var readme2nft Readme // define user object to marshal into

	json.Unmarshal(jsonObj, &readme2nft) // convert json string into the user object
	readme2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", readme2nft.Key, readme2nft.NftJSON)
	assert.Equal(t, expected, readme2nft.NftJSON, "check nft json against expected results")

	var nft2readme Readme // define user object to marshal into

	nft2readme.NftJSON = expected     // set the nft json
	nft2readme.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2readme.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2readme.NftJSON, "check unmarshalled against expected results")

}
