package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		  "_key": "bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi",
		  "name": "admin",
		  "domain": {
			"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
			"name": "GLOBAL"
		  },
		  "email": "admin@ortelius.io",
		  "phone": "505-444-5566",
		  "realname": "Ortelius Admin"
	  }`)

	expected := `{"domain":{"_key":"bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty"},"email":"admin@ortelius.io","name":"admin","objtype":"User","phone":"505-444-5566","realname":"Ortelius Admin"}`

	var user2nft User // define user object to marshal into

	json.Unmarshal(jsonObj, &user2nft)       // convert json string into the user object
	nftJSON := user2nft.MarshalNFT(cid2json) // generate the cid and nft json for user object
	assert.Equal(t, expected, nftJSON, "check nft json against expected results")

	var nft2user User // define user object to marshal into

	nft2user.Key = user2nft.Key            // set the nft json
	nft2user.UnmarshalNFT(cid2json)        // convert the json string into the user object
	check := nft2user.MarshalNFT(cid2json) // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, check, "check unmarshalled user against expected results")

}
