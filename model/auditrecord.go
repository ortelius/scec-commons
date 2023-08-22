// Package model - AuditRecord defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"time"
)

// AuditRecord defines a single audit event
type AuditRecord struct {
	Key     string    `json:"_key,omitempty"`
	ObjType string    `json:"objtype,omitempty"`
	Action  string    `json:"action"`
	User    *User     `json:"User"`
	When    time.Time `json:"when"`
}

// NewAuditRecord is the contructor that sets the appropriate default values
func NewAuditRecord() *AuditRecord {
	return &AuditRecord{ObjType: "AuditRecord", User: NewUser()}
}
