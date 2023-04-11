package model

import "encoding/json"

type Packages struct {
	Key      string    `json:"_key,omitempty"`
	NftJSON  string    `json:"_json,omitempty"`
	Packages []Package `json:"packages,omitempty"`
}

func (obj *Packages) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	type PackageNFT struct {
		Packages []NFT `json:"packages,omitempty"`
	}
	var packlist PackageNFT

	for _, v := range obj.Packages {
		nft := new(NFT).Init(v.MarshalNFT(cid2json))
		packlist.Packages = append(packlist.Packages, nft)
	}

	data, _ := json.Marshal(packlist)
	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *Packages) UnmarshalNFT(cid2json map[string]string) {
	var pkgs Packages // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &pkgs) // Convert the nft json into the domain object

	// Deep Copy
	obj.Packages = make([]Package, 0)

	for _, v := range pkgs.Packages {
		var rec Package

		rec.Key = v.Key
		rec.UnmarshalNFT(cid2json)
		obj.Packages = append(obj.Packages, rec)
	}
}
