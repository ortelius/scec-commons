// Package model - Group defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Group defines a group of users
type Group struct {
	Key    string `json:"_key,omitempty"`
	Domain Domain `json:"domain"`
	Name   string `json:"name"`
}
