package model

import "encoding/json"

type Package struct {
	Key        string `json:"_key,omitempty"`
	NftJson    string `json:"_json,omitempty"`
	License    string `json:"license,omitempty"`
	LicenseKey string `json:"license_key,omitempty"`
	Name       string `json:"name"`
	Purl       string `json:"purl,omitempty"`
	Version    string `json:"version"`
}

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

	obj.NftJson = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJson // Add cid=json for persisting later

	return data
}

func (obj *Package) UnmarshalNFT(cid2json map[string]string) {
	var pkg Package // define domain object to marshal into
	var exists bool
	var NftJson string

	// get the json from storage
	if NftJson, exists = cid2json[obj.Key]; exists {
		obj.NftJson = NftJson // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJson), &pkg) // Convert the nft json into the domain object

	// Deep Copy
	obj.License = pkg.License
	obj.LicenseKey = pkg.LicenseKey
	obj.Name = pkg.Name
	obj.Purl = pkg.Purl
	obj.Version = pkg.Version

}
