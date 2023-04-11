package model

import "encoding/json"

type DeploymentDetails struct {
	Key        string     `json:"_key,omitempty"`
	NftJSON    string     `json:"_json,omitempty"`
	Deployment Deployment `json:"deployment"`
	Log        []string   `json:"log,omitempty"`
}

func (obj *DeploymentDetails) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Deployment NFT      `json:"deployment"`
		Log        []string `json:"log,omitempty"`
		ObjType    string   `json:"objtype"`
	}{
		Deployment: new(NFT).Init(obj.Deployment.MarshalNFT(cid2json)),
		Log:        obj.Log,
		ObjType:    "DeploymentDetails",
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *DeploymentDetails) UnmarshalNFT(cid2json map[string]string) {
	var deploydetails DeploymentDetails
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &deploydetails)

	// Deep Copy
	obj.Deployment.Key = deploydetails.Deployment.Key
	obj.Deployment.UnmarshalNFT(cid2json)

	obj.Log = append(obj.Log, deploydetails.Log...)
}
