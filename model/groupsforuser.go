// Package model - GroupsForUser defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// GroupsForUser defines a user to a group
type GroupsForUser struct {
	Key       string   `json:"_key,omitempty"`
	GroupKeys []string `json:"groups"`
	UserKey   string   `json:"user"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *GroupsForUser) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		GroupKeys []string `json:"groups"`
		ObjType   string   `json:"objtype"`
		UserKey   string   `json:"user"`
	}{
		GroupKeys: obj.GroupKeys,
		ObjType:   "GroupsForUser",
		UserKey:   obj.UserKey,
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *GroupsForUser) UnmarshalNFT(cid2json map[string]string) {
	var groups4user GroupsForUser
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &groups4user)

		if err == nil {
			// Deep Copy
			obj.UserKey = groups4user.UserKey
			obj.GroupKeys = append(obj.GroupKeys, groups4user.GroupKeys...)
		}
	}
}
