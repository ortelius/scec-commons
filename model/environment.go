// Package model - Environment defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"encoding/json"
	"time"
)

// Environment defines a logical location that the deployment was perform against
type Environment struct {
	Key     string    `json:"_key,omitempty"`
	Created time.Time `json:"created"`
	Creator User      `json:"creator"`
	Domain  Domain    `json:"domain"`
	Name    string    `json:"name"`
	Owner   User      `json:"owner"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Environment) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Created time.Time `json:"created"`
		Creator NFT       `json:"creator"`
		Domain  NFT       `json:"domain"`
		Name    string    `json:"name"`
		ObjType string    `json:"objtype"`
		Owner   NFT       `json:"owner"`
	}{
		Created: obj.Created,
		Creator: new(NFT).Init(obj.Creator.MarshalNFT(cid2json)),
		Domain:  new(NFT).Init(obj.Domain.MarshalNFT(cid2json)),
		Name:    obj.Name,
		ObjType: "Environment",
		Owner:   new(NFT).Init(obj.Owner.MarshalNFT(cid2json)),
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Environment) UnmarshalNFT(cid2json map[string]string) {
	var environment Environment
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &environment)

		if err == nil {
			// Deep Copy
			obj.Created = environment.Created
			obj.Creator.Key = environment.Creator.Key
			obj.Creator.UnmarshalNFT(cid2json)
			obj.Domain.Key = environment.Domain.Key
			obj.Domain.UnmarshalNFT(cid2json)
			obj.Name = environment.Name
			obj.Owner.Key = environment.Owner.Key
			obj.Owner.UnmarshalNFT(cid2json)
		}
	}
}
