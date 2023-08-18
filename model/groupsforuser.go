// Package model - GroupsForUser defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// GroupsForUser defines a user to a group
type GroupsForUser struct {
	Key       string   `json:"_key,omitempty"`
	GroupKeys []string `json:"groups"`
	UserKey   string   `json:"user"`
}
