package model

import "encoding/json"

type Consuming struct {
	Key      string   `json:"_key,omitempty"`
	NftJSON  string   `json:"_json,omitempty"`
	Comsumes []string `json:"consumes"`
}

func (obj *Consuming) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Comsumes []string `json:"consumes"`
		ObjType  string   `json:"objtype"`
	}{
		Comsumes: obj.Comsumes,
		ObjType:  "Consuming",
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *Consuming) UnmarshalNFT(cid2json map[string]string) {
	var consuming Consuming
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &consuming)

	// Deep Copy
	obj.Comsumes = append(obj.Comsumes, consuming.Comsumes...)
}
