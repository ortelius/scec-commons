// Package model - Environment defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"time"
)

// Environment defines a logical location that the deployment was perform against
type Environment struct {
	Key     string    `json:"_key,omitempty"`
	Created time.Time `json:"created"`
	Creator User      `json:"creator"`
	Domain  Domain    `json:"domain"`
	Name    string    `json:"name"`
	Owner   User      `json:"owner"`
}
