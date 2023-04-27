// Package model - Packages defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Packages defines a list of Package
type Packages struct {
	Key      string    `json:"_key,omitempty"`
	Packages []Package `json:"packages,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Packages) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	type PackageNFT struct {
		Packages []NFT `json:"packages,omitempty"`
	}
	var packlist PackageNFT

	for _, v := range obj.Packages {
		nft := new(NFT).Init(v.MarshalNFT(cid2json))
		packlist.Packages = append(packlist.Packages, nft)
	}

	data, _ := json.Marshal(packlist)

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Packages) UnmarshalNFT(cid2json map[string]string) {
	var pkgs Packages // define domain object to marshal into
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &pkgs) // Convert the nft json into the domain object

		if err == nil {
			// Deep Copy
			obj.Packages = make([]Package, 0)

			for _, v := range pkgs.Packages {
				var rec Package

				rec.Key = v.Key
				rec.UnmarshalNFT(cid2json)
				obj.Packages = append(obj.Packages, rec)
			}
		}
	}
}
