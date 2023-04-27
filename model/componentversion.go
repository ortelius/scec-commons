// Package model - ComponentVersion defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// ComponentVersion defines a Version of an Component for a List View
type ComponentVersion struct {
	Key            string `json:"_key,omitempty"`
	Domain         Domain `json:"domain"`
	Name           string `json:"name"`
	ParentKey      string `json:"parent_key,omitempty"`
	PredecessorKey string `json:"predecessor_key,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *ComponentVersion) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Domain         NFT    `json:"domain"`
		Name           string `json:"name"`
		ObjType        string `json:"objtype"`
		ParentKey      string `json:"parent_key,omitempty"`
		PredecessorKey string `json:"predecessor_key,omitempty"`
	}{
		Domain:         new(NFT).Init(obj.Domain.MarshalNFT(cid2json)),
		Name:           obj.Name,
		ObjType:        "ComponentVersion",
		ParentKey:      obj.ParentKey,
		PredecessorKey: obj.PredecessorKey,
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *ComponentVersion) UnmarshalNFT(cid2json map[string]string) {
	var compver ComponentVersion // define domain object to marshal into
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {
		err := json.Unmarshal([]byte(nftJSON), &compver) // Convert the nft json into the domain object

		if err == nil {
			// Deep Copy
			obj.Domain.Key = compver.Domain.Key
			obj.Domain.UnmarshalNFT(cid2json)

			obj.Name = compver.Name
			obj.ParentKey = compver.ParentKey
			obj.PredecessorKey = compver.PredecessorKey
		}
	}
}
