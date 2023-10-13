// Package model - Package defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// Package defines a SBOM package dependency
type Package struct {
	Key     string `json:"_key,omitempty"`
	ObjType string `json:"objtype,omitempty"`
	CVE     string `json:"cve"`
	License string `json:"license,omitempty"`
	Name    string `json:"name"`
	Purl    string `json:"purl,omitempty"`
	Summary string `json:"summary"`
	Version string `json:"version"`
}

// NewPackage is the contructor that sets the appropriate default values
func NewPackage() *Package {
	return &Package{ObjType: "Package"}
}
