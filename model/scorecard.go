// Package model - Scorecard defines the struct and handles the OSSF Scorecard for a repo + commit sha
package model

// Scorecard defines a OSSF scorecard for a repo + commit.
// This struct is a collapsed version of the data returned from the api.scorecard.dev endpoint
type Scorecard struct {
	CommitSha            string  `json:"commit_sha"`
	Pinned               bool    `json:"pinned"`
	Score                float32 `json:"score"`
	Maintained           float32 `json:"maintained"`
	CodeReview           float32 `json:"code_review"`
	CIIBestPractices     float32 `json:"cii_best_practices"`
	License              float32 `json:"license"`
	SignedReleases       float32 `json:"signed_releases"`
	DangerousWorkflow    float32 `json:"dangerous_workflow"`
	Packaging            float32 `json:"packaging"`
	TokenPermissions     float32 `json:"token_permissions"`
	BranchProtection     float32 `json:"branch_protection"`
	BinaryArtifacts      float32 `json:"binary_artifacts"`
	PinnedDependencies   float32 `json:"pinned_dependencies"`
	SecurityPolicy       float32 `json:"security_policy"`
	Fuzzing              float32 `json:"fuzzing"`
	SAST                 float32 `json:"sast"`
	Vulnerabilities      float32 `json:"vulnerabilities"`
	CITests              float32 `json:"ci_tests"`
	Contributors         float32 `json:"contributors"`
	DependencyUpdateTool float32 `json:"dependency_update_tool"`
	SBOM                 float32 `json:"sbom"`
	Webhooks             float32 `json:"webhooks"`
}

// NewScorecard is the contructor that sets the appropriate default values
func NewScorecard() *Scorecard {
	return &Scorecard{}
}
