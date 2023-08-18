// Package model - Readme defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Readme defines a readme markdown file
type Readme struct {
	Key     string   `json:"_key,omitempty"`
	Content []string `json:"content"`
}
