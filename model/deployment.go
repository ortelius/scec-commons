// Package model - Deployment defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"time"
)

// Deployment defines a deployment for a List View
type Deployment struct {
	Key         string              `json:"_key,omitempty"`
	ObjType     string              `json:"objtype,omitempty"`
	Application *ApplicationVersion `json:"application"`
	Components  *Components         `json:"components"`
	DeployNum   int                 `json:"deploynum"`
	EndTime     time.Time           `json:"endtime,omitempty"`
	Environment *Environment        `json:"environment"`
	Result      int                 `json:"result,omitempty"`
	StartTime   time.Time           `json:"starttime"`
}

func NewDeployment() *Deployment {
	return &Deployment{ObjType: "Deployment", Application: NewApplicationVersion(), Components: NewComponents(), Environment: NewEnvironment()}
}
