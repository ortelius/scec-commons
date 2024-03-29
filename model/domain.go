// Package model - Domain defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Domain defines a dotted domain hierarchy
type Domain struct {
	Key     string `json:"_key,omitempty"`
	ObjType string `json:"objtype,omitempty"`
	Name    string `json:"name"`
}

// NewDomain is the contructor that sets the appropriate default values
func NewDomain() *Domain {
	return &Domain{ObjType: "Domain"}
}
