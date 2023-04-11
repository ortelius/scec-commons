package model

import "encoding/json"

type Providing struct {
	Key      string   `json:"_key,omitempty"`
	NftJSON  string   `json:"_json,omitempty"`
	Provides []string `json:"provides"`
}

func (obj *Providing) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		ObjType  string   `json:"objtype"`
		Provides []string `json:"provides"`
	}{
		ObjType:  "Providing",
		Provides: obj.Provides,
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *Providing) UnmarshalNFT(cid2json map[string]string) {
	var providing Providing
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &providing)

	// Deep Copy
	obj.Provides = append(obj.Provides, providing.Provides...)
}
