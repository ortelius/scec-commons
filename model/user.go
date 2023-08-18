// Package model - User defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// User defines a user
type User struct {
	Key      string `json:"_key,omitempty"`
	Name     string `json:"name"`
	Domain   Domain `json:"domain"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Realname string `json:"realname,omitempty"`
}
