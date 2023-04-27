package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsuming(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreibz4duaceeggbwnl7zhvzqbttyekglpwzac4dr57ig37fzwjvdcaa",
		"consumes": ["/user"]
	  }`)

	expected := `{"consumes":["/user"],"objtype":"Consuming"}`

	var consuming2nft Consuming // define user object to marshal into

	json.Unmarshal(jsonObj, &consuming2nft)       // convert json string into the user object
	nftJSON := consuming2nft.MarshalNFT(cid2json) // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", consuming2nft.Key, consuming2nft.NftJSON)
	assert.Equal(t, expected, nftJSON, "check nft json against expected results")

	var nft2consuming Consuming // define user object to marshal into

	nft2consuming.Key = consuming2nft.Key       // set the nft json
	nft2consuming.UnmarshalNFT(cid2json)        // convert the json string into the user object
	check := nft2consuming.MarshalNFT(cid2json) // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, check, "check unmarshalled user against expected results")

}
