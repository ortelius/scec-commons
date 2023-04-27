// Package model - Consuming defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Consuming defines a list of RestAPI end points being consumed by the Component Version
type Consuming struct {
	Key      string   `json:"_key,omitempty"`
	Comsumes []string `json:"consumes"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Consuming) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Comsumes []string `json:"consumes"`
		ObjType  string   `json:"objtype"`
	}{
		Comsumes: obj.Comsumes,
		ObjType:  "Consuming",
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Consuming) UnmarshalNFT(cid2json map[string]string) {
	var consuming Consuming
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &consuming)

		if err == nil {
			// Deep Copy
			obj.Comsumes = append(obj.Comsumes, consuming.Comsumes...)
		}
	}
}
