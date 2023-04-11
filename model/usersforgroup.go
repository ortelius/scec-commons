package model

import "encoding/json"

type UsersForGroup struct {
	Key       string   `json:"_key,omitempty"`
	NftJson   string   `json:"_json,omitempty"`
	Group_Key string   `json:"group"`
	User_Keys []string `json:"users"`
}

func (obj *UsersForGroup) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Group_Key string   `json:"group"`
		ObjType   string   `json:"objtype"`
		User_Keys []string `json:"users"`
	}{
		Group_Key: obj.Group_Key,
		ObjType:   "UsersForGroup",
		User_Keys: obj.User_Keys,
	})

	obj.NftJson = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJson // Add cid=json for persisting later

	return data
}

func (obj *UsersForGroup) UnmarshalNFT(cid2json map[string]string) {
	var users4group UsersForGroup
	var exists bool
	var NftJson string

	// get the json from storage
	if NftJson, exists = cid2json[obj.Key]; exists {
		obj.NftJson = NftJson // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJson), &users4group)

	// Deep Copy
	obj.Group_Key = users4group.Group_Key
	obj.User_Keys = append(obj.User_Keys, users4group.User_Keys...)
}
