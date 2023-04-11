package model

import (
	"encoding/json"
	"time"
)

type ApplicationVersionDetails struct {
	Key             string     `json:"_key,omitempty"`
	NftJson         string     `json:"_json,omitempty"`
	AuditLog        AuditLog   `json:"auditlog,omitempty"`
	Components      Components `json:"components,omitempty"`
	Created         time.Time  `json:"created"`
	Creator         User       `json:"creator"`
	Domain          Domain     `json:"domain"`
	Name            string     `json:"name"`
	Owner           User       `json:"owner"`
	Parent_Key      string     `json:"parent_key,omitempty"`
	Predecessor_Key string     `json:"predecessor_key,omitempty"`
}

func (obj *ApplicationVersionDetails) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		AuditLog        NFT       `json:"auditlog,omitempty"`
		Components      NFT       `json:"components,omitempty"`
		Created         time.Time `json:"created"`
		Creator         NFT       `json:"creator"`
		Domain          NFT       `json:"domain"`
		Name            string    `json:"name"`
		ObjType         string    `json:"objtype"`
		Owner           NFT       `json:"owner"`
		Parent_Key      string    `json:"parent_key,omitempty"`
		Predecessor_Key string    `json:"predecessor_key,omitempty"`
	}{
		AuditLog:        new(NFT).Init(obj.AuditLog.MarshalNFT(cid2json)),
		Components:      new(NFT).Init(obj.Components.MarshalNFT(cid2json)),
		Created:         obj.Created,
		Creator:         new(NFT).Init(obj.Creator.MarshalNFT(cid2json)),
		Domain:          new(NFT).Init(obj.Domain.MarshalNFT(cid2json)),
		Name:            obj.Name,
		ObjType:         "ApplicationVersionDetails",
		Owner:           new(NFT).Init(obj.Owner.MarshalNFT(cid2json)),
		Parent_Key:      obj.Parent_Key,
		Predecessor_Key: obj.Predecessor_Key,
	})

	obj.NftJson = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJson // Add cid=json for persisting later

	return data
}

func (obj *ApplicationVersionDetails) UnmarshalNFT(cid2json map[string]string) {
	var appver ApplicationVersionDetails // define domain object to marshal into
	var exists bool
	var NftJson string

	// get the json from storage
	if NftJson, exists = cid2json[obj.Key]; exists {
		obj.NftJson = NftJson // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJson), &appver) // Convert the nft json into the domain object

	// Deep Copy
	obj.AuditLog.Key = appver.AuditLog.Key
	obj.AuditLog.UnmarshalNFT(cid2json)

	obj.Components.Key = appver.Components.Key
	obj.Components.UnmarshalNFT(cid2json)

	obj.Created = appver.Created

	obj.Creator.Key = appver.Creator.Key
	obj.Creator.UnmarshalNFT(cid2json)

	obj.Domain.Key = appver.Domain.Key
	obj.Domain.UnmarshalNFT(cid2json)

	obj.Name = appver.Name

	obj.Owner.Key = appver.Owner.Key
	obj.Owner.UnmarshalNFT(cid2json)

	obj.Parent_Key = appver.Parent_Key
	obj.Predecessor_Key = appver.Predecessor_Key
}
