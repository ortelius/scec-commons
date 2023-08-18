// Package model - Vulnerabilities defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Vulnerabilities defines a list of Vulnerability
type Vulnerabilities struct {
	Key             string          `json:"_key,omitempty"`
	Vulnerabilities []Vulnerability `json:"vulnerabilties,omitempty"`
}
