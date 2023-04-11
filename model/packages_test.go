package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackages(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreia455fosaucpob7sebue7ae5pojh2gjrc5fijytq5ruunhn7ndfpi",
		"packages": [{
				"_key": "bafkreie72z3l77p6nkpkrmfyxqxopwnjiq3ztkur7ayhloleyqljsuf5ve",
				"purl": "pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
				"name": "libc-bin",
				"version": "2.19.18+deb8u7",
				"license_key": 23,
				"license": "GP-2.0"
			},
			{
				"_key": "bafkreianebpkdcvcna7ewjmpcspbw7k67lpf6oiuawdizxuwv6gnojrdla",
				"purl": "pkg:deb/debian/libcpp-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
				"name": "libcpp-bin",
				"version": "2.19.18+deb8u7",
				"license_key": 23,
				"license": "GP-2.0"
			}
		]
	}`)

	expected := `{"packages":[{"_key":"bafkreie72z3l77p6nkpkrmfyxqxopwnjiq3ztkur7ayhloleyqljsuf5ve"},{"_key":"bafkreianebpkdcvcna7ewjmpcspbw7k67lpf6oiuawdizxuwv6gnojrdla"}]}`

	var pkgs2nft Packages // define user object to marshal into

	json.Unmarshal(jsonObj, &pkgs2nft) // convert json string into the user object
	pkgs2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", pkgs2nft.Key, pkgs2nft.NftJson)
	assert.Equal(t, expected, pkgs2nft.NftJson, "check nft json against expected results")

	var nft2pkgs Packages // define user object to marshal into

	nft2pkgs.NftJson = expected     // set the nft json
	nft2pkgs.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2pkgs.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2pkgs.NftJson, "check unmarshalled user against expected results")

}
