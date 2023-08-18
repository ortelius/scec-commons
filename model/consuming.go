// Package model - Consuming defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Consuming defines a list of RestAPI end points being consumed by the Component Version
type Consuming struct {
	Key      string   `json:"_key,omitempty"`
	Comsumes []string `json:"consumes"`
}
