// Package model - License defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// License defines a license file for a Component Version
type License struct {
	Key     string   `json:"_key,omitempty"`
	ObjType string   `json:"objtype,omitempty"`
	Content []string `json:"content"`
}

// NewLicense is the contructor that sets the appropriate default values
func NewLicense() *License {
	return &License{ObjType: "License"}
}
