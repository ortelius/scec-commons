package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviding(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreifpq5fvyuajskoik4j7n362edrr6ubkoxj5gfm74gk2lxdpsw2fmy",
		"provides": ["/user"]
	}`)

	expected := `{"objtype":"Providing","provides":["/user"]}`

	var providing2nft Providing // define user object to marshal into

	json.Unmarshal(jsonObj, &providing2nft) // convert json string into the user object
	providing2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", providing2nft.Key, providing2nft.NftJSON)
	assert.Equal(t, expected, providing2nft.NftJSON, "check nft json against expected results")

	var nft2providing Providing // define user object to marshal into

	nft2providing.NftJSON = expected     // set the nft json
	nft2providing.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2providing.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2providing.NftJSON, "check unmarshalled user against expected results")

}
