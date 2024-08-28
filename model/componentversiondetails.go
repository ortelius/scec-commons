// Package model - ComponentVersionDetails defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"time"
)

// Note: Graphs are used to connect the License, Readme and Swagger to the Component Version
// since those objects more than likely be redundant across compvers.  So we will only store them once
// and reference multiple times via the graph relationships.

// ComponentVersionDetails defines a Version of a Component including fine grained details
type ComponentVersionDetails struct {
	Key       string     `json:"_key,omitempty"`
	ObjType   string     `json:"objtype,omitempty"`
	Attrs     *CompAttrs `json:"attrs,omitempty"`
	AuditLog  *AuditLog  `json:"autditlog,omitempty"`
	CompType  string     `json:"comptype,omitempty"`
	Created   time.Time  `json:"created,omitempty"`
	Creator   *User      `json:"creator,omitempty"`
	Domain    *Domain    `json:"domain,omitempty"`
	License   *License   `json:"license,omitempty"` // In `license` collection via comp2lic graph
	Name      string     `json:"name,omitempty"`
	Owner     *User      `json:"owner,omitempty"`
	Packages  []*Package `json:"packages,omitempty"`  // SBOMs are stored using the compid so we can get them directly
	Readme    *Readme    `json:"readme,omitempty"`    // In `readme` collection via comp2readme graph
	Scorecard *Scorecard `json:"scorecard,omitempty"` // Scorecards are stored using the compid so we can get them directly
	Swagger   *Swagger   `json:"swagger,omitempty"`   // In `swagger` collection via comp2swagger graph
	Variant   string     `json:"variant,omitempty"`
	Version   string     `json:"version,omitempty"`
}

// NewComponentVersionDetails is the contructor that sets the appropriate default values
func NewComponentVersionDetails() *ComponentVersionDetails {
	return &ComponentVersionDetails{
		ObjType:   "ComponentVersionDetails",
		Attrs:     NewCompAttrs(),
		AuditLog:  NewAuditLog(),
		Creator:   NewUser(),
		Domain:    NewDomain(),
		License:   NewLicense(),
		Owner:     NewUser(),
		Packages:  NewPackages(),
		Readme:    NewReadme(),
		Scorecard: NewScorecard(),
		Swagger:   NewSwagger()}
}
