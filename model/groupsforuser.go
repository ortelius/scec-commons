package model

import "encoding/json"

type GroupsForUser struct {
	Key        string   `json:"_key,omitempty"`
	NftJson    string   `json:"_json,omitempty"`
	Group_Keys []string `json:"groups"`
	User_Key   string   `json:"user"`
}

func (obj *GroupsForUser) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Group_Keys []string `json:"groups"`
		ObjType    string   `json:"objtype"`
		User_Key   string   `json:"user"`
	}{
		Group_Keys: obj.Group_Keys,
		ObjType:    "GroupsForUser",
		User_Key:   obj.User_Key,
	})

	obj.NftJson = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJson // Add cid=json for persisting later

	return data
}

func (obj *GroupsForUser) UnmarshalNFT(cid2json map[string]string) {
	var groups4user GroupsForUser
	var exists bool
	var NftJson string

	// get the json from storage
	if NftJson, exists = cid2json[obj.Key]; exists {
		obj.NftJson = NftJson // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJson), &groups4user)

	// Deep Copy
	obj.User_Key = groups4user.User_Key
	obj.Group_Keys = append(obj.Group_Keys, groups4user.Group_Keys...)
}
