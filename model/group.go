// Package model - Group defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Group defines a group of users
type Group struct {
	Key     string  `json:"_key,omitempty"`
	ObjType string  `json:"objtype,omitempty"`
	Domain  *Domain `json:"domain"`
	Name    string  `json:"name"`
}

// NewGroup is the contructor that sets the appropriate default values
func NewGroup() *Group {
	return &Group{ObjType: "Group", Domain: NewDomain()}
}
