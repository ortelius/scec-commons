// Package ortelius - ApplicationVersion defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package ortelius

import "encoding/json"

// ApplicationVersion defines a Version of an Application for a List View
type ApplicationVersion struct {
	Key            string `json:"_key,omitempty"`
	NftJSON        string `json:"_json,omitempty"`
	Deployments    []int  `json:"deployments,omitempty"`
	Domain         Domain `json:"domain"`
	Name           string `json:"name"`
	ParentKey      string `json:"parent_key,omitempty"`
	PredecessorKey string `json:"predecessor_key,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *ApplicationVersion) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Deployments    []int  `json:"deployments,omitempty"`
		Domain         NFT    `json:"domain"`
		Name           string `json:"name"`
		ObjType        string `json:"objtype"`
		ParentKey      string `json:"parent_key,omitempty"`
		PredecessorKey string `json:"predecessor_key,omitempty"`
	}{
		Deployments:    obj.Deployments,
		Domain:         new(NFT).Init(obj.Domain.MarshalNFT(cid2json)),
		Name:           obj.Name,
		ObjType:        "ApplicationVersion",
		ParentKey:      obj.ParentKey,
		PredecessorKey: obj.PredecessorKey,
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *ApplicationVersion) UnmarshalNFT(cid2json map[string]string) {
	var appver ApplicationVersion // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	err := json.Unmarshal([]byte(obj.NftJSON), &appver) // Convert the nft json into the domain object

	if err == nil {
		// Deep Copy
		obj.Domain.Key = appver.Domain.Key
		obj.Domain.UnmarshalNFT(cid2json)

		obj.Name = appver.Name
		obj.ParentKey = appver.ParentKey
		obj.PredecessorKey = appver.PredecessorKey
		obj.Deployments = append(obj.Deployments, appver.Deployments...)
	}
}
