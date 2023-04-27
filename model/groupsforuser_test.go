package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupsForUser(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreic3fzlubuzkgfly7glgx7apgwnmt6kdgiasqr4gc3ycayj43nmqme",
		"user": "admin",
		"groups": [ "users", "administrators" ]
	  }`)

	expected := `{"groups":["users","administrators"],"objtype":"GroupsForUser","user":"admin"}`

	var grps4usr2nft GroupsForUser // define user object to marshal into

	json.Unmarshal(jsonObj, &grps4usr2nft)       // convert json string into the user object
	nftJSON := grps4usr2nft.MarshalNFT(cid2json) // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", grps4usr2nft.Key, grps4usr2nft.NftJSON)
	assert.Equal(t, expected, nftJSON, "check nft json against expected results")

	var nft2grps4usr GroupsForUser // define user object to marshal into

	nft2grps4usr.Key = grps4usr2nft.Key        // set the nft json
	nft2grps4usr.UnmarshalNFT(cid2json)        // convert the json string into the user object
	check := nft2grps4usr.MarshalNFT(cid2json) // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, check, "check unmarshalled against expected results")

}
