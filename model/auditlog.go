// Package model - AuditLog defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// AuditLog defines a list of Audit Records
type AuditLog struct {
	Key      string        `json:"_key,omitempty"`
	AuditLog []AuditRecord `json:"auditlog,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *AuditLog) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	type AuditLogNFT struct {
		AuditLog []NFT `json:"auditlog,omitempty"`
	}
	var reclist AuditLogNFT

	for _, v := range obj.AuditLog {
		nft := new(NFT).Init(v.MarshalNFT(cid2json))
		reclist.AuditLog = append(reclist.AuditLog, nft)
	}

	data, _ := json.Marshal(reclist)

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *AuditLog) UnmarshalNFT(cid2json map[string]string) {
	var auditlog AuditLog // define domain object to marshal into
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &auditlog) // Convert the nft json into the domain object

		if err == nil {
			// Deep Copy
			obj.AuditLog = make([]AuditRecord, 0)

			for _, v := range auditlog.AuditLog {
				var rec AuditRecord

				rec.Key = v.Key
				rec.UnmarshalNFT(cid2json)
				obj.AuditLog = append(obj.AuditLog, rec)
			}
		}
	}
}
