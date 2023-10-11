// Package model - ResponseKey defines the struct to return the NFT key as a json object to the client
package model

// ResponseKey defines a struct to hold the NFT key in a json response
type ResponseKey struct {
	Key string `json:"_key,omitempty"`
}
