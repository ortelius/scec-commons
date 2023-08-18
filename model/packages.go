// Package model - Packages defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Packages defines a list of Package
type Packages struct {
	Key      string    `json:"_key,omitempty"`
	Packages []Package `json:"packages,omitempty"`
}
