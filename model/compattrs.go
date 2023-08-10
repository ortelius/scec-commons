// Package model - CompAttrs defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"encoding/json"
	"time"
)

// CompAttrs defines the well known attributes for a Component
type CompAttrs struct {
	Key                      string    `json:"_key,omitempty"`
	Basename                 string    `json:"basename,omitempty"`
	BuildDate                time.Time `json:"builddate,omitempty"`
	BuildID                  string    `json:"buildid,omitempty"`
	BuildNum                 string    `json:"buildnum,omitempty"`
	BuildURL                 string    `json:"buildurl,omitempty"`
	Chart                    string    `json:"chart,omitempty"`
	ChartNamespace           string    `json:"chartnamespace,omitempty"`
	ChartRepo                string    `json:"chartrepo,omitempty"`
	ChartRepoURL             string    `json:"chartrepourl,omitempty"`
	ChartVersion             string    `json:"chartversion,omitempty"`
	DiscordChannel           string    `json:"discordchannel,omitempty"`
	DockerRepo               string    `json:"dockerrepo,omitempty"`
	DockerSha                string    `json:"dockersha,omitempty"`
	DockerTag                string    `json:"dockertag,omitempty"`
	GitBranch                string    `json:"gitbranch,omitempty"`
	GitBranchCreateCommit    string    `json:"gitbranchcreatecommit,omitempty"`
	GitBranchCreateTimestamp time.Time `json:"gitbranchcreatetimestamp,omitempty"`
	GitBranchParent          string    `json:"gitbranchparent,omitempty"`
	GitCommit                string    `json:"gitcommit,omitempty"`
	GitCommitAuthors         string    `json:"gitcommitauthors,omitempty"`
	GitCommittersCnt         string    `json:"gitcommittescnt,omitempty"`
	GitCommitTimestamp       time.Time `json:"gitcommittimestamp,omitempty"`
	GitContribPercentage     string    `json:"gitcontribpercentage,omitempty"`
	GitLinesAdded            string    `json:"gitlinesadded,omitempty"`
	GitLinesDeleted          string    `json:"gitlinesdeleted,omitempty"`
	GitLinesTotal            string    `json:"gitlinestotal,omitempty"`
	GitOrg                   string    `json:"gitorg,omitempty"`
	GitPrevCompCommit        string    `json:"gitpreviouscomponentcommit,omitempty"`
	GitRepo                  string    `json:"gitrepo,omitempty"`
	GitRepoProject           string    `json:"gitrepoproject,omitempty"`
	GitSignedOffBy           string    `json:"gitsignedoffby,omitempty"`
	GitTag                   string    `json:"gittag,omitempty"`
	GitTotalCommittersCnt    string    `json:"gittotalcommittescnt,omitempty"`
	GitURL                   string    `json:"giturl,omitempty"`
	GitVerifyCommit          bool      `json:"gitverifycommit,omitempty"`
	HipchatChannel           string    `json:"hipchatchannel,omitempty"`
	PagerdutyBusinessURL     string    `json:"pagerdutybusinessurl,omitempty"`
	PagerdutyURL             string    `json:"pagerdutyurl,omitempty"`
	Repository               string    `json:"repository,omitempty"`
	ServiceOwner             User      `json:"serviceowner,omitempty"`
	SlackChannel             string    `json:"slackchannel,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *CompAttrs) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Basename                 string    `json:"basename,omitempty"`
		BuildDate                time.Time `json:"builddate,omitempty"`
		BuildID                  string    `json:"buildid,omitempty"`
		BuildNum                 string    `json:"buildnum,omitempty"`
		BuildURL                 string    `json:"buildurl,omitempty"`
		Chart                    string    `json:"chart,omitempty"`
		ChartNamespace           string    `json:"chartnamespace,omitempty"`
		ChartRepo                string    `json:"chartrepo,omitempty"`
		ChartRepoURL             string    `json:"chartrepourl,omitempty"`
		ChartVersion             string    `json:"chartversion,omitempty"`
		DiscordChannel           string    `json:"discordchannel,omitempty"`
		DockerRepo               string    `json:"dockerrepo,omitempty"`
		DockerSha                string    `json:"dockersha,omitempty"`
		DockerTag                string    `json:"dockertag,omitempty"`
		GitBranch                string    `json:"gitbranch,omitempty"`
		GitBranchCreateCommit    string    `json:"gitbranchcreatecommit,omitempty"`
		GitBranchCreateTimestamp time.Time `json:"gitbranchcreatetimestamp,omitempty"`
		GitBranchParent          string    `json:"gitbranchparent,omitempty"`
		GitCommit                string    `json:"gitcommit,omitempty"`
		GitCommitAuthors         string    `json:"gitcommitauthors,omitempty"`
		GitCommittersCnt         string    `json:"gitcommittescnt,omitempty"`
		GitCommitTimestamp       time.Time `json:"gitcommittimestamp,omitempty"`
		GitContribPercentage     string    `json:"gitcontribpercentage,omitempty"`
		GitLinesAdded            string    `json:"gitlinesadded,omitempty"`
		GitLinesDeleted          string    `json:"gitlinesdeleted,omitempty"`
		GitLinesTotal            string    `json:"gitlinestotal,omitempty"`
		GitOrg                   string    `json:"gitorg,omitempty"`
		GitPrevCompCommit        string    `json:"gitpreviouscomponentcommit,omitempty"`
		GitRepo                  string    `json:"gitrepo,omitempty"`
		GitRepoProject           string    `json:"gitrepoproject,omitempty"`
		GitSignedOffBy           string    `json:"gitsignedoffby,omitempty"`
		GitTag                   string    `json:"gittag,omitempty"`
		GitTotalCommittersCnt    string    `json:"gittotalcommittescnt,omitempty"`
		GitURL                   string    `json:"giturl,omitempty"`
		GitVerifyCommit          bool      `json:"gitverifycommit,omitempty"`
		HipchatChannel           string    `json:"hipchatchannel,omitempty"`
		ObjType                  string    `json:"objtype"`
		PagerdutyBusinessURL     string    `json:"pagerdutybusinessurl,omitempty"`
		PagerdutyURL             string    `json:"pagerdutyurl,omitempty"`
		Repository               string    `json:"repository,omitempty"`
		ServiceOwner             NFT       `json:"serviceowner,omitempty"`
		SlackChannel             string    `json:"slackchannel,omitempty"`
	}{
		Basename:                 obj.Basename,
		BuildDate:                obj.BuildDate,
		BuildID:                  obj.BuildID,
		BuildURL:                 obj.BuildURL,
		Chart:                    obj.Chart,
		ChartNamespace:           obj.ChartNamespace,
		ChartRepo:                obj.ChartRepo,
		ChartRepoURL:             obj.ChartRepoURL,
		ChartVersion:             obj.ChartVersion,
		DiscordChannel:           obj.DiscordChannel,
		DockerRepo:               obj.DockerRepo,
		DockerSha:                obj.DockerSha,
		DockerTag:                obj.DockerTag,
		GitBranch:                obj.GitBranch,
		GitBranchCreateCommit:    obj.GitBranchCreateCommit,
		GitBranchCreateTimestamp: obj.GitBranchCreateTimestamp,
		GitBranchParent:          obj.GitBranchParent,
		GitCommit:                obj.GitCommit,
		GitCommitAuthors:         obj.GitCommitAuthors,
		GitCommittersCnt:         obj.GitCommittersCnt,
		GitCommitTimestamp:       obj.GitCommitTimestamp,
		GitContribPercentage:     obj.GitContribPercentage,
		GitLinesAdded:            obj.GitLinesAdded,
		GitLinesDeleted:          obj.GitLinesDeleted,
		GitLinesTotal:            obj.GitLinesTotal,
		GitOrg:                   obj.GitOrg,
		GitPrevCompCommit:        obj.GitPrevCompCommit,
		GitRepo:                  obj.GitRepo,
		GitRepoProject:           obj.GitRepoProject,
		GitSignedOffBy:           obj.GitSignedOffBy,
		GitTag:                   obj.GitTag,
		GitTotalCommittersCnt:    obj.GitTotalCommittersCnt,
		GitURL:                   obj.GitURL,
		GitVerifyCommit:          obj.GitVerifyCommit,
		HipchatChannel:           obj.HipchatChannel,
		ObjType:                  "CompAttr",
		PagerdutyBusinessURL:     obj.PagerdutyBusinessURL,
		PagerdutyURL:             obj.PagerdutyURL,
		Repository:               obj.Repository,
		ServiceOwner:             new(NFT).Init(obj.ServiceOwner.MarshalNFT(cid2json)),
		SlackChannel:             obj.SlackChannel,
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *CompAttrs) UnmarshalNFT(cid2json map[string]string) {
	var compattrs CompAttrs
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &compattrs)

		if err == nil {
			// Deep Copy
			obj.Basename = compattrs.Basename
			obj.BuildDate = compattrs.BuildDate
			obj.BuildID = compattrs.BuildID
			obj.BuildURL = compattrs.BuildURL
			obj.Chart = compattrs.Chart
			obj.ChartNamespace = compattrs.ChartNamespace
			obj.ChartRepo = compattrs.ChartRepo
			obj.ChartRepoURL = compattrs.ChartRepoURL
			obj.ChartVersion = compattrs.ChartVersion
			obj.DiscordChannel = compattrs.DiscordChannel
			obj.DockerRepo = compattrs.DockerRepo
			obj.DockerSha = compattrs.DockerSha
			obj.DockerTag = compattrs.DockerTag
			obj.GitBranch = compattrs.GitBranch
			obj.GitBranchCreateCommit = compattrs.GitBranchCreateCommit
			obj.GitBranchCreateTimestamp = compattrs.GitBranchCreateTimestamp
			obj.GitBranchParent = compattrs.GitBranchParent
			obj.GitCommit = compattrs.GitCommit
			obj.GitCommitAuthors = compattrs.GitCommitAuthors
			obj.GitCommittersCnt = compattrs.GitCommittersCnt
			obj.GitCommitTimestamp = compattrs.GitCommitTimestamp
			obj.GitContribPercentage = compattrs.GitContribPercentage
			obj.GitLinesAdded = compattrs.GitLinesAdded
			obj.GitLinesDeleted = compattrs.GitLinesDeleted
			obj.GitLinesTotal = compattrs.GitLinesTotal
			obj.GitOrg = compattrs.GitOrg
			obj.GitPrevCompCommit = compattrs.GitPrevCompCommit
			obj.GitRepo = compattrs.GitRepo
			obj.GitRepoProject = compattrs.GitRepoProject
			obj.GitSignedOffBy = compattrs.GitSignedOffBy
			obj.GitTag = compattrs.GitTag
			obj.GitTotalCommittersCnt = compattrs.GitTotalCommittersCnt
			obj.GitURL = compattrs.GitURL
			obj.GitVerifyCommit = compattrs.GitVerifyCommit
			obj.HipchatChannel = compattrs.HipchatChannel
			obj.PagerdutyBusinessURL = compattrs.PagerdutyBusinessURL
			obj.PagerdutyURL = compattrs.PagerdutyURL
			obj.Repository = compattrs.Repository
			obj.ServiceOwner.Key = compattrs.ServiceOwner.Key
			obj.ServiceOwner.UnmarshalNFT(cid2json)
			obj.SlackChannel = compattrs.SlackChannel
		}
	}
}
