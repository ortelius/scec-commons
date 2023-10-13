// Package model - ComponentVersionDetails defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"time"
)

// ComponentVersionDetails defines a Version of a Component including fine grained details
type ComponentVersionDetails struct {
	Key            string     `json:"_key,omitempty"`
	ObjType        string     `json:"objtype,omitempty"`
	Attrs          *CompAttrs `json:"attrs,omitempty"`
	AuditLog       *AuditLog  `json:"autditlog,omitempty"`
	CompType       string     `json:"comptype"`
	Consuming      *Consuming `json:"consuming,omitempty"`
	Created        time.Time  `json:"created"`
	Creator        *User      `json:"creator"`
	Domain         *Domain    `json:"domain"`
	License        *License   `json:"license,omitempty"`
	Name           string     `json:"name"`
	Owner          *User      `json:"owner"`
	Packages       *Packages  `json:"packages,omitempty"`
	ParentKey      string     `json:"parent_key,omitempty"`
	PredecessorKey string     `json:"predecessor_key,omitempty"`
	Providing      *Providing `json:"providing,omitempty"`
	Readme         *Readme    `json:"readme,omitempty"`
	SBOMKey        string     `json:"sbom_key,omitempty"`
	Swagger        *Swagger   `json:"swagger,omitempty"`
}

// NewComponentVersionDetails is the contructor that sets the appropriate default values
func NewComponentVersionDetails() *ComponentVersionDetails {
	return &ComponentVersionDetails{
		ObjType:   "ComponentVersionDetails",
		Attrs:     NewCompAttrs(),
		AuditLog:  NewAuditLog(),
		Consuming: NewConsuming(),
		Creator:   NewUser(),
		Domain:    NewDomain(),
		License:   NewLicense(),
		Owner:     NewUser(),
		Packages:  NewPackages(),
		Providing: NewProviding(),
		Readme:    NewReadme(),
		Swagger:   NewSwagger()}
}
