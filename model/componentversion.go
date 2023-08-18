// Package model - ComponentVersion defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// ComponentVersion defines a Version of an Component for a List View
type ComponentVersion struct {
	Key            string `json:"_key,omitempty"`
	ObjType        string `default:"ComponentVersion" json:"objtype,omitempty"`
	Domain         Domain `json:"domain"`
	Name           string `json:"name"`
	ParentKey      string `json:"parent_key,omitempty"`
	PredecessorKey string `json:"predecessor_key,omitempty"`
}
