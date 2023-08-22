// Package model - Swagger defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Swagger defines an OpenAPI or Swagger file
type Swagger struct {
	Key     string          `json:"_key,omitempty"`
	ObjType string          `json:"objtype,omitempty"`
	Content json.RawMessage `json:"content"`
}

// NewSwagger is the contructor that sets the appropriate default values
func NewSwagger() *Swagger {
	return &Swagger{ObjType: "Swagger"}
}
