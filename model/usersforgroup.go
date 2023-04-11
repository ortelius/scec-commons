package model

import "encoding/json"

type UsersForGroup struct {
	Key      string   `json:"_key,omitempty"`
	NftJSON  string   `json:"_json,omitempty"`
	GroupKey string   `json:"group"`
	UserKeys []string `json:"users"`
}

func (obj *UsersForGroup) MarshalNFT(cid2json map[string]string) []byte {

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

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *UsersForGroup) UnmarshalNFT(cid2json map[string]string) {
	var users4group UsersForGroup
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &users4group)

	// Deep Copy
	obj.GroupKey = users4group.GroupKey
	obj.UserKeys = append(obj.UserKeys, users4group.UserKeys...)
}
