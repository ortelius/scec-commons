// Package model - AuditLog defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// AuditLog defines a list of Audit Records
type AuditLog struct {
	Key      string         `json:"_key,omitempty"`
	ObjType  string         `json:"objtype,omitempty"`
	AuditLog []*AuditRecord `json:"auditlog,omitempty"`
}

func NewAuditLog() *AuditLog {
	return &AuditLog{ObjType: "AuditLog"}
}
