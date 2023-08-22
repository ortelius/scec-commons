// Package model - Applications defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Applications defines a list of Application Versions
type Applications struct {
	Key          string                `json:"_key,omitempty"`
	ObjType      string                `json:"objtype,omitempty"`
	Applications []*ApplicationVersion `json:"applications,omitempty"`
}

func NewApplications() *Applications {
	return &Applications{ObjType: "Applications"}
}
