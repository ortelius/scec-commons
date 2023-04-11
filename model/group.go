package model

import "encoding/json"

type Group struct {
	Key     string `json:"_key,omitempty"`
	NftJSON string `json:"_json,omitempty"`
	Domain  Domain `json:"domain"`
	Name    string `json:"name"`
}

func (obj *Group) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Domain  NFT    `json:"domain"`
		Name    string `json:"name"`
		ObjType string `json:"objtype"`
	}{
		Domain:  new(NFT).Init(obj.Domain.MarshalNFT(cid2json)),
		Name:    obj.Name,
		ObjType: "Group",
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *Group) UnmarshalNFT(cid2json map[string]string) {
	var group Group // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &group) // Convert the nft json into the domain object

	// Deep Copy
	obj.Name = group.Name
	obj.Domain.Key = group.Domain.Key
	obj.Domain.UnmarshalNFT(cid2json)
}
