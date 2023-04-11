package model

import "encoding/json"

type Vulnerabilities struct {
	Key             string          `json:"_key,omitempty"`
	NftJSON         string          `json:"_json,omitempty"`
	Vulnerabilities []Vulnerability `json:"vulnerabilties,omitempty"`
}

func (obj *Vulnerabilities) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	type VulnerabilityNFT struct {
		Vulnerabilities []NFT `json:"vulnerabilties,omitempty"`
	}
	var vulnlist VulnerabilityNFT

	for _, v := range obj.Vulnerabilities {
		nft := new(NFT).Init(v.MarshalNFT(cid2json))
		vulnlist.Vulnerabilities = append(vulnlist.Vulnerabilities, nft)
	}

	data, _ := json.Marshal(vulnlist)
	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *Vulnerabilities) UnmarshalNFT(cid2json map[string]string) {
	var pkgs Vulnerabilities // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &pkgs) // Convert the nft json into the domain object

	// Deep Copy
	obj.Vulnerabilities = make([]Vulnerability, 0)

	for _, v := range pkgs.Vulnerabilities {
		var rec Vulnerability

		rec.Key = v.Key
		rec.UnmarshalNFT(cid2json)
		obj.Vulnerabilities = append(obj.Vulnerabilities, rec)
	}
}
