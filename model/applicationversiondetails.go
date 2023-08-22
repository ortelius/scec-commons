// Package model - ApplicationVersionDetails defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"time"
)

// ApplicationVersionDetails defines a Version of an Application including fine grained details
type ApplicationVersionDetails struct {
	Key            string      `json:"_key,omitempty"`
	ObjType        string      `json:"objtype,omitempty"`
	AuditLog       *AuditLog   `json:"auditlog,omitempty"`
	Components     *Components `json:"components,omitempty"`
	Created        time.Time   `json:"created"`
	Creator        *User       `json:"creator"`
	Domain         *Domain     `json:"domain"`
	Name           string      `json:"name"`
	Owner          *User       `json:"owner"`
	ParentKey      string      `json:"parent_key,omitempty"`
	PredecessorKey string      `json:"predecessor_key,omitempty"`
}

func NewApplicationVersionDetails() *ApplicationVersionDetails {
	return &ApplicationVersionDetails{ObjType: "ApplicationVersionDetails", AuditLog: NewAuditLog(), Components: NewComponents(), Creator: NewUser(), Domain: NewDomain(), Owner: NewUser()}
}
