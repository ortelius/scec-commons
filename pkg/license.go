// Package pkg - License defines the struct and handles marshalling/unmarshalling the struct to/from NFT Storage.
package pkg

import "encoding/json"

// License defines a license file for a Component Version
type License struct {
	Key     string   `json:"_key,omitempty"`
	NftJSON string   `json:"_json,omitempty"`
	Content []string `json:"content"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *License) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Content []string `json:"content"`
		ObjType string   `json:"objtype"`
	}{
		Content: obj.Content,
		ObjType: "License",
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *License) UnmarshalNFT(cid2json map[string]string) {
	var license License // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &license) // Convert the nft json into the domain object

	// Deep Copy
	obj.Content = append(obj.Content, license.Content...)
}