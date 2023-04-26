// Package model - Domain defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Domain defines a dotted domain hierarchy
type Domain struct {
	Key     string `json:"_key,omitempty"`
	NftJSON string `json:"_json,omitempty"`
	Name    string `json:"name"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Domain) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		ObjType string `json:"objtype"`
		Name    string `json:"name"`
	}{
		ObjType: "Domain", // Set
		Name:    obj.Name, // Copy
	})

	obj.NftJSON = string(data)        // Save the json
	obj.Key = new(NFT).Init(data).Key // Calculate and save the cid for the json
	cid2json[obj.Key] = obj.NftJSON   // Add cid=json for persisting later

	return data // Return NFT Json
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Domain) UnmarshalNFT(cid2json map[string]string) {
	var domain Domain // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	err := json.Unmarshal([]byte(obj.NftJSON), &domain) // Convert the nft json into the domain object

	if err == nil {
		// Deep Copy
		obj.Name = domain.Name
	}
}
