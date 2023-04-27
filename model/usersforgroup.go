// Package model - UsersForGroup defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// UsersForGroup defines a list of user that belong to the group
type UsersForGroup struct {
	Key      string   `json:"_key,omitempty"`
	GroupKey string   `json:"group"`
	UserKeys []string `json:"users"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *UsersForGroup) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		GroupKey string   `json:"group"`
		ObjType  string   `json:"objtype"`
		UserKeys []string `json:"users"`
	}{
		GroupKey: obj.GroupKey,
		ObjType:  "UsersForGroup",
		UserKeys: obj.UserKeys,
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *UsersForGroup) UnmarshalNFT(cid2json map[string]string) {
	var users4group UsersForGroup
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &users4group)

		if err == nil {
			// Deep Copy
			obj.GroupKey = users4group.GroupKey
			obj.UserKeys = append(obj.UserKeys, users4group.UserKeys...)
		}
	}
}
