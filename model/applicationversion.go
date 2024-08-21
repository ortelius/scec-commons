// Package model - ApplicationVersion defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// ApplicationVersion defines a Version of an Application for a List View
type ApplicationVersion struct {
	Key         string  `json:"_key,omitempty"`
	ObjType     string  `json:"objtype,omitempty"`
	Deployments []int   `json:"deployments,omitempty"`
	Domain      *Domain `json:"domain"`
	Name        string  `json:"name"`
	Variant     string  `json:"variant"`
	Version     string  `json:"version"`
}

// NewApplicationVersion is the contructor that sets the appropriate default values
func NewApplicationVersion() *ApplicationVersion {
	return &ApplicationVersion{ObjType: "ApplicationVersion", Domain: NewDomain()}
}
