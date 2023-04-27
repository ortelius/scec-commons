// Package model - Providing defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Providing defines a list of RestAPI endpoints exposed by the Component Version
type Providing struct {
	Key      string   `json:"_key,omitempty"`
	Provides []string `json:"provides"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Providing) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		ObjType  string   `json:"objtype"`
		Provides []string `json:"provides"`
	}{
		ObjType:  "Providing",
		Provides: obj.Provides,
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Providing) UnmarshalNFT(cid2json map[string]string) {
	var providing Providing
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &providing)

		if err == nil {
			// Deep Copy
			obj.Provides = append(obj.Provides, providing.Provides...)
		}
	}
}
