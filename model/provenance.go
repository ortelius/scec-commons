// Package model - SBOM defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Provenance defines the Provenance of an component in JSON format
type Provenance struct {
	Key     string          `json:"_key,omitempty"`
	ObjType string          `json:"objtype,omitempty"`
	Content json.RawMessage `json:"content"`
}

// NewProvenance is the contructor that sets the appropriate default values
func NewProvenance() *Provenance {
	return &Provenance{ObjType: "Provenance"}
}
