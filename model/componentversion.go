// Package model - ComponentVersion defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// ComponentVersion defines a Version of an Component for a List View
type ComponentVersion struct {
	Key     string  `json:"_key,omitempty"`
	ObjType string  `json:"objtype,omitempty"`
	Domain  *Domain `json:"domain"`
	Name    string  `json:"name"`
	Variant string  `json:"variant"`
	Version string  `json:"version"`
}

// NewComponentVersion is the contructor that sets the appropriate default values
func NewComponentVersion() *ComponentVersion {
	return &ComponentVersion{ObjType: "ComponentVersion", Domain: NewDomain()}
}
