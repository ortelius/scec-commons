// Package model - DeploymentDetails defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// DeploymentDetails defines a deployment plus the associated log
type DeploymentDetails struct {
	Key        string      `json:"_key,omitempty"`
	ObjType    string      `json:"objtype,omitempty"`
	Deployment *Deployment `json:"deployment"`
	Log        []string    `json:"log,omitempty"`
}

// NewDeploymentDetails is the contructor that sets the appropriate default values
func NewDeploymentDetails() *DeploymentDetails {
	return &DeploymentDetails{ObjType: "DeploymentDetails", Deployment: NewDeployment()}
}
