package pkg

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomain(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
		"name": "GLOBAL"
	  }`)

	expected := `{"objtype":"Domain","name":"GLOBAL"}`

	var dom2nft Domain // define user object to marshal into

	json.Unmarshal(jsonObj, &dom2nft) // convert json string into the user object
	dom2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", dom2nft.Key, dom2nft.NftJSON)
	assert.Equal(t, expected, dom2nft.NftJSON, "check nft json against expected results")

	var nft2dom Domain // define user object to marshal into

	nft2dom.NftJSON = expected     // set the nft json
	nft2dom.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2dom.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2dom.NftJSON, "check unmarshalled against expected results")

}
