// Package model - Package defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Package defines a SBOM package dependency
type Package struct {
	Key        string `json:"_key,omitempty"`
	NftJSON    string `json:"_json,omitempty"`
	License    string `json:"license,omitempty"`
	LicenseKey string `json:"license_key,omitempty"`
	Name       string `json:"name"`
	Purl       string `json:"purl,omitempty"`
	Version    string `json:"version"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Package) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		License    string `json:"license,omitempty"`
		LicenseKey string `json:"license_key,omitempty"`
		Name       string `json:"name"`
		ObjType    string `json:"objtype"`
		Purl       string `json:"purl,omitempty"`
		Version    string `json:"version"`
	}{
		License:    obj.License,
		LicenseKey: obj.LicenseKey,
		Name:       obj.Name,
		ObjType:    "Package",
		Purl:       obj.Purl,
		Version:    obj.Version,
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Package) UnmarshalNFT(cid2json map[string]string) {
	var pkg Package // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	err := json.Unmarshal([]byte(obj.NftJSON), &pkg) // Convert the nft json into the domain object

	if err == nil {
		// Deep Copy
		obj.License = pkg.License
		obj.LicenseKey = pkg.LicenseKey
		obj.Name = pkg.Name
		obj.Purl = pkg.Purl
		obj.Version = pkg.Version
	}

}
