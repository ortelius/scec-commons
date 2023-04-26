package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLicense(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreicv5dl7ozitvglmfm6jsqvw3f2pqqwpocgl6vbefgratj2gohtvmy",
		"content": ["# Apache 2", "## Summary"]
	  }`)

	expected := `{"content":["# Apache 2","## Summary"],"objtype":"License"}`

	var license2nft License // define user object to marshal into

	json.Unmarshal(jsonObj, &license2nft) // convert json string into the user object
	license2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", license2nft.Key, license2nft.NftJSON)
	assert.Equal(t, expected, license2nft.NftJSON, "check nft json against expected results")

	var nft2license License // define user object to marshal into

	nft2license.NftJSON = expected     // set the nft json
	nft2license.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2license.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2license.NftJSON, "check unmarshalled against expected results")

}
