// Package model - License defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// License defines a license file for a Component Version
type License struct {
	Key     string   `json:"_key,omitempty"`
	Content []string `json:"content"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *License) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Content []string `json:"content"`
		ObjType string   `json:"objtype"`
	}{
		Content: obj.Content,
		ObjType: "License",
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *License) UnmarshalNFT(cid2json map[string]string) {
	var license License // define domain object to marshal into
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &license) // Convert the nft json into the domain object

		if err == nil {
			// Deep Copy
			obj.Content = append(obj.Content, license.Content...)
		}
	}
}
