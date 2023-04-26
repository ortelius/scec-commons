// Package ortelius - GroupsForUser defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package ortelius

import "encoding/json"

// GroupsForUser defines a user to a group
type GroupsForUser struct {
	Key       string   `json:"_key,omitempty"`
	NftJSON   string   `json:"_json,omitempty"`
	GroupKeys []string `json:"groups"`
	UserKey   string   `json:"user"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *GroupsForUser) MarshalNFT(cid2json map[string]string) []byte {

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

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *GroupsForUser) UnmarshalNFT(cid2json map[string]string) {
	var groups4user GroupsForUser
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	err := json.Unmarshal([]byte(obj.NftJSON), &groups4user)

	if err == nil {
		// Deep Copy
		obj.UserKey = groups4user.UserKey
		obj.GroupKeys = append(obj.GroupKeys, groups4user.GroupKeys...)
	}
}
