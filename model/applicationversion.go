// Package model - ApplicationVersion defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// ApplicationVersion defines a Version of an Application for a List View
type ApplicationVersion struct {
	Key            string  `json:"_key,omitempty"`
	ObjType        string  `json:"objtype,omitempty"`
	Deployments    []int   `json:"deployments,omitempty"`
	Domain         *Domain `json:"domain"`
	Name           string  `json:"name"`
	ParentKey      string  `json:"parent_key,omitempty"`
	PredecessorKey string  `json:"predecessor_key,omitempty"`
}

func NewApplicationVersion() *ApplicationVersion {
	return &ApplicationVersion{ObjType: "ApplicationVersion", Domain: NewDomain()}
}
