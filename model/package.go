// Package model - Package defines the struct for returning a subset of the SBOM to the UI for rendering
package model

// Package defines a SBOM package dependency (subset of the full sbom)
type Package struct {
	CVE      string `json:"cve"`
	Language string `json:"language"`
	License  string `json:"license"`
	Name     string `json:"name"`
	Purl     string `json:"purl"`
	Summary  string `json:"summary"`
	Version  string `json:"version"`
}

// NewPackage is the contructor that sets the appropriate default values
func NewPackage() *Package {
	return &Package{}
}
