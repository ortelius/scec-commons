package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwagger(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreihntxpcz47l7blasyee3tv2as6qthk44yifi7blwswi6w5dbnpcza",
		"content": ["# Rest APIs", "## GET /user"]
	  }`)

	expected := `{"content":["# Rest APIs","## GET /user"],"objtype":"Swagger"}`

	var swagger2nft Swagger // define user object to marshal into

	json.Unmarshal(jsonObj, &swagger2nft) // convert json string into the user object
	swagger2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", swagger2nft.Key, swagger2nft.NftJson)
	assert.Equal(t, expected, swagger2nft.NftJson, "check nft json against expected results")

	var nft2swagger Swagger // define user object to marshal into

	nft2swagger.NftJson = expected     // set the nft json
	nft2swagger.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2swagger.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2swagger.NftJson, "check unmarshalled against expected results")

}
