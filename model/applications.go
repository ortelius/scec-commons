// Package model - Applications defines the struct and handles marshalling/unmarshalling the struct to/from NFT Storage.
package model

import "encoding/json"

// Applications defines a list of Application Versions
type Applications struct {
	Key          string               `json:"_key,omitempty"`
	NftJSON      string               `json:"_json,omitempty"`
	Applications []ApplicationVersion `json:"applications,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Applications) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	type ApplicationVersionNFT struct {
		Applications []NFT `json:"applications,omitempty"`
	}
	var applist ApplicationVersionNFT

	for _, v := range obj.Applications {
		nft := new(NFT).Init(v.MarshalNFT(cid2json))
		applist.Applications = append(applist.Applications, nft)
	}

	data, _ := json.Marshal(applist)
	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Applications) UnmarshalNFT(cid2json map[string]string) {
	var apps Applications // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &apps) // Convert the nft json into the domain object

	// Deep Copy
	obj.Applications = make([]ApplicationVersion, 0)

	for _, v := range apps.Applications {
		var rec ApplicationVersion

		rec.Key = v.Key
		rec.UnmarshalNFT(cid2json)
		obj.Applications = append(obj.Applications, rec)
	}
}
