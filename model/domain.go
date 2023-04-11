package model

import "encoding/json"

type Domain struct {
	Key     string `json:"_key,omitempty"`
	NftJson string `json:"_json,omitempty"`
	Name    string `json:"name"`
}

func (obj *Domain) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		ObjType string `json:"objtype"`
		Name    string `json:"name"`
	}{
		ObjType: "Domain", // Set
		Name:    obj.Name, // Copy
	})

	obj.NftJson = string(data)        // Save the json
	obj.Key = new(NFT).Init(data).Key // Calculate and save the cid for the json
	cid2json[obj.Key] = obj.NftJson   // Add cid=json for persisting later

	return data // Return NFT Json
}

func (obj *Domain) UnmarshalNFT(cid2json map[string]string) {
	var domain Domain // define domain object to marshal into
	var exists bool
	var NftJson string

	// get the json from storage
	if NftJson, exists = cid2json[obj.Key]; exists {
		obj.NftJson = NftJson // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJson), &domain) // Convert the nft json into the domain object

	// Deep Copy
	obj.Name = domain.Name
}
