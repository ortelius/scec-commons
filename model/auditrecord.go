// Package model - AuditRecord defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"encoding/json"
	"time"
)

// AuditRecord defines a single audit event
type AuditRecord struct {
	Key    string    `json:"_key,omitempty"`
	Action string    `json:"action"`
	User   User      `json:"User"`
	When   time.Time `json:"when"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *AuditRecord) MarshalNFT(cid2json map[string]string) string {

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

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *AuditRecord) UnmarshalNFT(cid2json map[string]string) {
	var auditrecord AuditRecord
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &auditrecord)

		if err == nil {
			// Deep Copy
			obj.Action = auditrecord.Action
			obj.When = auditrecord.When
			obj.User.Key = auditrecord.User.Key
			obj.User.UnmarshalNFT(cid2json)
		}
	}
}
