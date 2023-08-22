// Package model - Components defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Components defines a list of Component Versions
type Components struct {
	Key        string              `json:"_key,omitempty"`
	ObjType    string              `json:"objtype,omitempty"`
	Components []*ComponentVersion `json:"components,omitempty"`
}

// NewComponents is the contructor that sets the appropriate default values
func NewComponents() *Components {
	return &Components{ObjType: "Components"}
}
