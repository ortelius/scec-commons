// Package model - ComponentVersionDetails defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"encoding/json"
	"time"
)

// ComponentVersionDetails defines a Version of a Component including fine grained details
type ComponentVersionDetails struct {
	Key             string          `json:"_key,omitempty"`
	Applications    Applications    `json:"applications,omitempty"`
	Attrs           CompAttrs       `json:"attrs,omitempty"`
	AuditLog        AuditLog        `json:"autditlog,omitempty"`
	CompType        string          `json:"comptype"`
	Consuming       Consuming       `json:"consuming,omitempty"`
	Created         time.Time       `json:"created"`
	Creator         User            `json:"creator"`
	Domain          Domain          `json:"domain"`
	License         License         `json:"license,omitempty"`
	Name            string          `json:"name"`
	Owner           User            `json:"owner"`
	Packages        Packages        `json:"packages,omitempty"`
	ParentKey       string          `json:"parent_key,omitempty"`
	PredecessorKey  string          `json:"predecessor_key,omitempty"`
	Providing       Providing       `json:"providing,omitempty"`
	Readme          Readme          `json:"readme,omitempty"`
	Swagger         Swagger         `json:"swagger,omitempty"`
	Vulnerabilities Vulnerabilities `json:"vulnerabilties,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *ComponentVersionDetails) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Applications    NFT       `json:"applications,omitempty"`
		Attrs           NFT       `json:"attrs,omitempty"`
		AuditLog        NFT       `json:"autditlog,omitempty"`
		CompType        string    `json:"comptype"`
		Consuming       NFT       `json:"consuming,omitempty"`
		Created         time.Time `json:"created"`
		Creator         NFT       `json:"creator"`
		Domain          NFT       `json:"domain"`
		License         NFT       `json:"license,omitempty"`
		Name            string    `json:"name"`
		ObjType         string    `json:"objtype"`
		Owner           NFT       `json:"owner"`
		Packages        NFT       `json:"packages,omitempty"`
		ParentKey       string    `json:"parent_key,omitempty"`
		PredecessorKey  string    `json:"predecessor_key,omitempty"`
		Providing       NFT       `json:"providing,omitempty"`
		Readme          NFT       `json:"readme,omitempty"`
		Swagger         NFT       `json:"swagger,omitempty"`
		Vulnerabilities NFT       `json:"vulnerabilities,omitempty"`
	}{
		Applications:    new(NFT).Init(obj.Applications.MarshalNFT(cid2json)),
		Attrs:           new(NFT).Init(obj.Attrs.MarshalNFT(cid2json)),
		AuditLog:        new(NFT).Init(obj.AuditLog.MarshalNFT(cid2json)),
		CompType:        obj.CompType,
		Consuming:       new(NFT).Init(obj.Consuming.MarshalNFT(cid2json)),
		Created:         obj.Created,
		Creator:         new(NFT).Init(obj.Creator.MarshalNFT(cid2json)),
		Domain:          new(NFT).Init(obj.Domain.MarshalNFT(cid2json)),
		License:         new(NFT).Init(obj.License.MarshalNFT(cid2json)),
		Name:            obj.Name,
		ObjType:         "ComponentVersionDetails",
		Owner:           new(NFT).Init(obj.Owner.MarshalNFT(cid2json)),
		Packages:        new(NFT).Init(obj.Packages.MarshalNFT(cid2json)),
		ParentKey:       obj.ParentKey,
		PredecessorKey:  obj.PredecessorKey,
		Providing:       new(NFT).Init(obj.Providing.MarshalNFT(cid2json)),
		Readme:          new(NFT).Init(obj.Readme.MarshalNFT(cid2json)),
		Swagger:         new(NFT).Init(obj.Swagger.MarshalNFT(cid2json)),
		Vulnerabilities: new(NFT).Init(obj.Vulnerabilities.MarshalNFT(cid2json)),
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *ComponentVersionDetails) UnmarshalNFT(cid2json map[string]string) {
	var compver ComponentVersionDetails // define domain object to marshal into
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &compver) // Convert the nft json into the domain object

		if err == nil {
			// Deep Copy
			//	obj.Applications.Key = compver.Applications.Key
			//	obj.Applications.UnmarshalNFT(cid2json)

			obj.Attrs.Key = compver.Attrs.Key
			obj.Attrs.UnmarshalNFT(cid2json)

			obj.AuditLog.Key = compver.AuditLog.Key
			obj.AuditLog.UnmarshalNFT(cid2json)

			obj.CompType = compver.CompType

			obj.Consuming.Key = compver.Consuming.Key
			obj.Consuming.UnmarshalNFT(cid2json)

			obj.Created = compver.Created

			obj.Creator.Key = compver.Creator.Key
			obj.Creator.UnmarshalNFT(cid2json)

			obj.Domain.Key = compver.Domain.Key
			obj.Domain.UnmarshalNFT(cid2json)

			obj.License.Key = compver.License.Key
			obj.License.UnmarshalNFT(cid2json)

			obj.Name = compver.Name

			obj.Owner.Key = compver.Owner.Key
			obj.Owner.UnmarshalNFT(cid2json)

			obj.Packages.Key = compver.Packages.Key
			obj.Packages.UnmarshalNFT(cid2json)

			obj.ParentKey = compver.ParentKey
			obj.PredecessorKey = compver.PredecessorKey

			obj.Providing.Key = compver.Providing.Key
			obj.Providing.UnmarshalNFT(cid2json)

			obj.Readme.Key = compver.Readme.Key
			obj.Readme.UnmarshalNFT(cid2json)

			obj.Swagger.Key = compver.Swagger.Key
			obj.Swagger.UnmarshalNFT(cid2json)

			obj.Vulnerabilities.Key = compver.Vulnerabilities.Key
			obj.Vulnerabilities.UnmarshalNFT(cid2json)
		}
	}
}
