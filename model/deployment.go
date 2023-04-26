// Package model - Deployment defines the struct and handles marshalling/unmarshalling the struct to/from NFT Storage.
package model

import (
	"encoding/json"
	"time"
)

// Deployment defines a deployment for a List View
type Deployment struct {
	Key         string             `json:"_key,omitempty"`
	NftJSON     string             `json:"_json,omitempty"`
	Application ApplicationVersion `json:"application"`
	Components  Components         `json:"components"`
	DeployNum   int                `json:"deploynum"`
	EndTime     time.Time          `json:"endtime,omitempty"`
	Environment Environment        `json:"environment"`
	Result      int                `json:"result,omitempty"`
	StartTime   time.Time          `json:"starttime"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Deployment) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Application NFT       `json:"application"`
		Components  NFT       `json:"components"`
		DeployNum   int       `json:"deploynum"`
		EndTime     time.Time `json:"endtime,omitempty"`
		Environment NFT       `json:"environment"`
		ObjType     string    `json:"objtype"`
		Result      int       `json:"result,omitempty"`
		StartTime   time.Time `json:"starttime"`
	}{
		Application: new(NFT).Init(obj.Application.MarshalNFT(cid2json)),
		Components:  new(NFT).Init(obj.Components.MarshalNFT(cid2json)),
		DeployNum:   obj.DeployNum,
		EndTime:     obj.EndTime,
		Environment: new(NFT).Init(obj.Environment.MarshalNFT(cid2json)),
		ObjType:     "Deployment",
		Result:      obj.Result,
		StartTime:   obj.StartTime,
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Deployment) UnmarshalNFT(cid2json map[string]string) {
	var deployment Deployment
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &deployment)

	// Deep Copy
	obj.Application.Key = deployment.Application.Key
	obj.Application.UnmarshalNFT(cid2json)

	obj.Components.Key = deployment.Components.Key
	obj.Components.UnmarshalNFT(cid2json)

	obj.Environment.Key = deployment.Environment.Key
	obj.Environment.UnmarshalNFT(cid2json)

	obj.DeployNum = deployment.DeployNum
	obj.EndTime = deployment.EndTime
	obj.Result = deployment.Result
	obj.StartTime = deployment.StartTime
}
