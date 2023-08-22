// Package model - Readme defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Readme defines a readme markdown file
type Readme struct {
	Key     string   `json:"_key,omitempty"`
	ObjType string   `json:"objtype,omitempty"`
	Content []string `json:"content"`
}

// NewReadme is the contructor that sets the appropriate default values
func NewReadme() *Readme {
	return &Readme{ObjType: "Readme"}
}
