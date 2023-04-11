package model

import (
	"encoding/json"
	"time"
)

type AuditRecord struct {
	Key     string    `json:"_key,omitempty"`
	NftJson string    `json:"_json,omitempty"`
	Action  string    `json:"action"`
	User    User      `json:"User"`
	When    time.Time `json:"when"`
}

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

	obj.NftJson = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJson // Add cid=json for persisting later

	return data
}

func (obj *AuditRecord) UnmarshalNFT(cid2json map[string]string) {
	var auditrecord AuditRecord
	var exists bool
	var NftJson string

	// get the json from storage
	if NftJson, exists = cid2json[obj.Key]; exists {
		obj.NftJson = NftJson // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJson), &auditrecord)

	// Deep Copy
	obj.Action = auditrecord.Action
	obj.When = auditrecord.When
	obj.User.Key = auditrecord.User.Key
	obj.User.UnmarshalNFT(cid2json)
}
