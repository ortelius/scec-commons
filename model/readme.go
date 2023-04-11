package model

import "encoding/json"

type Readme struct {
	Key     string   `json:"_key,omitempty"`
	NftJSON string   `json:"_json,omitempty"`
	Content []string `json:"content"`
}

func (obj *Readme) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Content []string `json:"content"`
		ObjType string   `json:"objtype"`
	}{
		Content: obj.Content,
		ObjType: "Readme",
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *Readme) UnmarshalNFT(cid2json map[string]string) {
	var readme Readme // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &readme) // Convert the nft json into the domain object

	// Deep Copy
	obj.Content = append(obj.Content, readme.Content...)
}
