// Package model - GroupsForUser defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// GroupsForUser defines a user to a group
type GroupsForUser struct {
	Key       string   `json:"_key,omitempty"`
	ObjType   string   `json:"objtype,omitempty"`
	GroupKeys []string `json:"groups"`
	UserKey   string   `json:"user"`
}

func NewGroupsForUser() *GroupsForUser {
	return &GroupsForUser{ObjType: "GroupsForUser"}
}
