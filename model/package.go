// Package model - Package defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Package defines a SBOM package dependency
type Package struct {
	Key        string `json:"_key,omitempty"`
	License    string `json:"license,omitempty"`
	LicenseKey string `json:"license_key,omitempty"`
	Name       string `json:"name"`
	Purl       string `json:"purl,omitempty"`
	Version    string `json:"version"`
}
