package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersForGroup(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreidoh27gg2wqnitspexsqjv2vjq56s5w2sgkeskbqqctgpjwntehkm",
		"group": "administrators",
		"users": [ "testadmin", "admin" ]
	  }`)

	expected := `{"group":"administrators","objtype":"UsersForGroup","users":["testadmin","admin"]}`

	var users4grp2nft UsersForGroup // define user object to marshal into

	json.Unmarshal(jsonObj, &users4grp2nft)       // convert json string into the user object
	nftJSON := users4grp2nft.MarshalNFT(cid2json) // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", users4grp2nft.Key, users4grp2nft.NftJSON)
	assert.Equal(t, expected, nftJSON, "check nft json against expected results")

	var nft2users4grp UsersForGroup // define user object to marshal into

	nft2users4grp.Key = users4grp2nft.Key       // set the nft json
	nft2users4grp.UnmarshalNFT(cid2json)        // convert the json string into the user object
	check := nft2users4grp.MarshalNFT(cid2json) // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, check, "check unmarshalled against expected results")

}
