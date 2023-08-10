package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwagger(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreiesors5xgoehdkpe6dc36y2eweosixks6qkpfi6qj36ed4o4c6ene",
		"content": {"openapi":"3.0.2"}
	  }`)

	expected := `{"content":{"openapi":"3.0.2"},"objtype":"Swagger"}`

	var swagger2nft Swagger // define user object to marshal into

	json.Unmarshal(jsonObj, &swagger2nft)       // convert json string into the user object
	nftJSON := swagger2nft.MarshalNFT(cid2json) // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", swagger2nft.Key, nftJSON)
	assert.Equal(t, expected, nftJSON, "check nft json against expected results")

	var nft2swagger Swagger // define user object to marshal into

	nft2swagger.Key = swagger2nft.Key         // set the nft json
	nft2swagger.UnmarshalNFT(cid2json)        // convert the json string into the user object
	check := nft2swagger.MarshalNFT(cid2json) // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, check, "check unmarshalled against expected results")

}
