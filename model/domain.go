// Package model - Domain defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Domain defines a dotted domain hierarchy
type Domain struct {
	Key  string `json:"_key,omitempty"`
	Name string `json:"name"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Domain) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		ObjType string `json:"objtype"`
		Name    string `json:"name"`
	}{
		ObjType: "Domain", // Set
		Name:    obj.Name, // Copy
	})

	obj.Key = new(NFT).Init(string(data)).Key // Calculate and save the cid for the json
	cid2json[obj.Key] = string(data)          // Add cid=json for persisting later

	return string(data) // Return NFT Json
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Domain) UnmarshalNFT(cid2json map[string]string) {
	var domain Domain // define domain object to marshal into
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &domain) // Convert the nft json into the domain object

		if err == nil {
			// Deep Copy
			obj.Name = domain.Name
		}
	}
}
