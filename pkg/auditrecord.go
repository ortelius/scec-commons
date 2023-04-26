// Package pkg - AuditRecord defines the struct and handles marshalling/unmarshalling the struct to/from NFT Storage.
package pkg

import (
	"encoding/json"
	"time"
)

// AuditRecord defines a single audit event
type AuditRecord struct {
	Key     string    `json:"_key,omitempty"`
	NftJSON string    `json:"_json,omitempty"`
	Action  string    `json:"action"`
	User    User      `json:"User"`
	When    time.Time `json:"when"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *AuditRecord) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Action  string    `json:"action"`
		ObjType string    `json:"objtype"`
		User    NFT       `json:"User"`
		When    time.Time `json:"when"`
	}{
		Action:  obj.Action,
		ObjType: "AuditRecord",
		User:    new(NFT).Init(obj.User.MarshalNFT(cid2json)),
		When:    obj.When,
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *AuditRecord) UnmarshalNFT(cid2json map[string]string) {
	var auditrecord AuditRecord
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &auditrecord)

	// Deep Copy
	obj.Action = auditrecord.Action
	obj.When = auditrecord.When
	obj.User.Key = auditrecord.User.Key
	obj.User.UnmarshalNFT(cid2json)
}
