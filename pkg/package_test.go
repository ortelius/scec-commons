package ortelius

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreie72z3l77p6nkpkrmfyxqxopwnjiq3ztkur7ayhloleyqljsuf5ve",
		"purl": "pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
		"name": "libc-bin",
		"version": "2.19.18+deb8u7",
		"license_key": 23,
		"license": "GP-2.0"
	  }`)

	expected := `{"license":"GP-2.0","name":"libc-bin","objtype":"Package","purl":"pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64\u0026upstream=glibc\u0026distro=debian-8","version":"2.19.18+deb8u7"}`

	var pkg2nft Package // define user object to marshal into

	json.Unmarshal(jsonObj, &pkg2nft) // convert json string into the user object
	pkg2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", pkg2nft.Key, pkg2nft.NftJSON)
	assert.Equal(t, expected, pkg2nft.NftJSON, "check nft json against expected results")

	var nft2pkg Package // define user object to marshal into

	nft2pkg.NftJSON = expected     // set the nft json
	nft2pkg.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2pkg.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2pkg.NftJSON, "check unmarshalled user against expected results")

}
