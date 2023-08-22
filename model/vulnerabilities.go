// Package model - Vulnerabilities defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Vulnerabilities defines a list of Vulnerability
type Vulnerabilities struct {
	Key             string           `json:"_key,omitempty"`
	ObjType         string           `json:"objtype,omitempty"`
	Vulnerabilities []*Vulnerability `json:"vulnerabilties,omitempty"`
}

// NewVulnerabilities is the contructor that sets the appropriate default values
func NewVulnerabilities() *Vulnerabilities {
	return &Vulnerabilities{ObjType: "Vulnerabilities"}
}
