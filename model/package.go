// Package model - Package defines the struct for returning a subset of the SBOM to the UI for rendering
package model

// Package defines a SBOM package dependency (subset of the full sbom)
// This struct is a collapsed version of the data returned from CycloneDX json struct
type Package struct {
	CVE      string  `json:"cve"`
	FixedIn  string  `json:"fixedin"`
	Language string  `json:"language"`
	License  string  `json:"license"`
	Name     string  `json:"name"`
	Purl     string  `json:"purl"`
	Score    float64 `json:"score"`
	Severity string  `json:"severity"`
	Summary  string  `json:"summary"`
	Version  string  `json:"version"`
}

// NewPackage is the contructor that sets the appropriate default values
func NewPackage() *Package {
	return &Package{}
}

// NewPackages is the constructor that returns an empty array of Package instances
func NewPackages() []*Package {
	return []*Package{}
}

// NOTE: PackageLicense and PackageCVE are for backward compatibility with v10 frontend

// PackageLicense defines a SBOM package dependency with the corresponding license (subset of the full sbom)
type PackageLicense struct {
	Key      string  `json:"key"`
	CompID   string  `json:"compid"`
	FixedIn  string  `json:"fixedin"`
	Language string  `json:"pkgtype"`
	License  string  `json:"name"` // name of the license
	Name     string  `json:"packagename"`
	Purl     string  `json:"purl"`
	Score    float64 `json:"score"`
	Severity string  `json:"risklevel"`
	Summary  string  `json:"summary"`
	URL      string  `json:"url"` // url to the license details
	Version  string  `json:"packageversion"`
}

// NewPackageLicense is the contructor that sets the appropriate default values
func NewPackageLicense() *PackageLicense {
	return &PackageLicense{}
}

// PackageCVE defines a SBOM package dependency with the corresponding license (subset of the full sbom)
type PackageCVE struct {
	Key      string  `json:"key"`
	CompID   string  `json:"compid"`
	CVE      string  `json:"name"` // CVE name
	FixedIn  string  `json:"fixedin"`
	Language string  `json:"pkgtype"`
	Name     string  `json:"packagename"`
	Purl     string  `json:"purl"`
	Score    float64 `json:"score"`
	Severity string  `json:"risklevel"`
	Summary  string  `json:"summary"`
	URL      string  `json:"url"` // CVE url on osv.dev
	Version  string  `json:"packageversion"`
}

// NewPackageCVE is the contructor that sets the appropriate default values
func NewPackageCVE() *PackageCVE {
	return &PackageCVE{}
}
