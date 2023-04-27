// Package model - Group defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Group defines a group of users
type Group struct {
	Key    string `json:"_key,omitempty"`
	Domain Domain `json:"domain"`
	Name   string `json:"name"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Group) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Domain  NFT    `json:"domain"`
		Name    string `json:"name"`
		ObjType string `json:"objtype"`
	}{
		Domain:  new(NFT).Init(obj.Domain.MarshalNFT(cid2json)),
		Name:    obj.Name,
		ObjType: "Group",
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Group) UnmarshalNFT(cid2json map[string]string) {
	var group Group // define domain object to marshal into
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &group) // Convert the nft json into the domain object

		if err == nil {
			// Deep Copy
			obj.Name = group.Name
			obj.Domain.Key = group.Domain.Key
			obj.Domain.UnmarshalNFT(cid2json)
		}
	}
}
