// Package model - Providing defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Providing defines a list of RestAPI endpoints exposed by the Component Version
type Providing struct {
	Key      string   `json:"_key,omitempty"`
	ObjType  string   `json:"objtype,omitempty"`
	Provides []string `json:"provides"`
}

// NewProviding is the contructor that sets the appropriate default values
func NewProviding() *Providing {
	return &Providing{ObjType: "Providing"}
}
