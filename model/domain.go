// Package model - Domain defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Domain defines a dotted domain hierarchy
type Domain struct {
	Name string `json:"name"`
}
