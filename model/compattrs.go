package model

import "encoding/json"

type CompAttrs struct {
	Key                  string `json:"_key,omitempty"`
	NftJSON              string `json:"_json,omitempty"`
	BuildDate            string `json:"builddate,omitempty"`
	BuildId              string `json:"buildid,omitempty"`
	BuildURL             string `json:"buildurl,omitempty"`
	Chart                string `json:"chart,omitempty"`
	ChartNamespace       string `json:"chartnamespace,omitempty"`
	ChartRepo            string `json:"chartrepo,omitempty"`
	ChartRepoURL         string `json:"chartrepourl,omitempty"`
	ChartVersion         string `json:"chartversion,omitempty"`
	DiscordChannel       string `json:"discordchannel,omitempty"`
	DockerRepo           string `json:"dockerrepo,omitempty"`
	DockerSha            string `json:"dockersha,omitempty"`
	DockerTag            string `json:"dockertag,omitempty"`
	GitCommit            string `json:"gitcommit,omitempty"`
	GitRepo              string `json:"gitrepo,omitempty"`
	GitTag               string `json:"gittag,omitempty"`
	GitURL               string `json:"giturl,omitempty"`
	HipchatChannel       string `json:"hipchatchannel,omitempty"`
	PagerdutyBusinessURL string `json:"pagerdutybusinessurl,omitempty"`
	PagerdutyURL         string `json:"pagerdutyurl,omitempty"`
	Repository           string `json:"repository,omitempty"`
	ServiceOwner         User   `json:"serviceowner,omitempty"`
	SlackChannel         string `json:"slackchannel,omitempty"`
}

func (obj *CompAttrs) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		BuildDate            string `json:"builddate,omitempty"`
		BuildId              string `json:"buildid,omitempty"`
		BuildURL             string `json:"buildurl,omitempty"`
		Chart                string `json:"chart,omitempty"`
		ChartNamespace       string `json:"chartnamespace,omitempty"`
		ChartRepo            string `json:"chartrepo,omitempty"`
		ChartRepoURL         string `json:"chartrepourl,omitempty"`
		ChartVersion         string `json:"chartversion,omitempty"`
		DiscordChannel       string `json:"discordchannel,omitempty"`
		DockerRepo           string `json:"dockerrepo,omitempty"`
		DockerSha            string `json:"dockersha,omitempty"`
		DockerTag            string `json:"dockertag,omitempty"`
		GitCommit            string `json:"gitcommit,omitempty"`
		GitRepo              string `json:"gitrepo,omitempty"`
		GitTag               string `json:"gittag,omitempty"`
		GitURL               string `json:"giturl,omitempty"`
		HipchatChannel       string `json:"hipchatchannel,omitempty"`
		ObjType              string `json:"objtype"`
		PagerdutyBusinessURL string `json:"pagerdutybusinessurl,omitempty"`
		PagerdutyURL         string `json:"pagerdutyurl,omitempty"`
		Repository           string `json:"repository,omitempty"`
		ServiceOwner         NFT    `json:"serviceowner,omitempty"`
		SlackChannel         string `json:"slackchannel,omitempty"`
	}{
		BuildDate:            obj.BuildDate,
		BuildId:              obj.BuildId,
		BuildURL:             obj.BuildURL,
		Chart:                obj.Chart,
		ChartNamespace:       obj.ChartNamespace,
		ChartRepo:            obj.ChartRepo,
		ChartRepoURL:         obj.ChartRepoURL,
		ChartVersion:         obj.ChartVersion,
		DiscordChannel:       obj.DiscordChannel,
		DockerRepo:           obj.DockerRepo,
		DockerSha:            obj.DockerSha,
		DockerTag:            obj.DockerTag,
		GitCommit:            obj.GitCommit,
		GitRepo:              obj.GitRepo,
		GitTag:               obj.GitTag,
		GitURL:               obj.GitURL,
		HipchatChannel:       obj.HipchatChannel,
		ObjType:              "CompAttr",
		PagerdutyBusinessURL: obj.PagerdutyBusinessURL,
		PagerdutyURL:         obj.PagerdutyURL,
		Repository:           obj.Repository,
		ServiceOwner:         new(NFT).Init(obj.ServiceOwner.MarshalNFT(cid2json)),
		SlackChannel:         obj.SlackChannel,
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *CompAttrs) UnmarshalNFT(cid2json map[string]string) {
	var compattrs CompAttrs
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &compattrs)

	// Deep Copy
	obj.BuildDate = compattrs.BuildDate
	obj.BuildId = compattrs.BuildId
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
	obj.GitCommit = compattrs.GitCommit
	obj.GitRepo = compattrs.GitRepo
	obj.GitTag = compattrs.GitTag
	obj.GitURL = compattrs.GitURL
	obj.HipchatChannel = compattrs.HipchatChannel
	obj.PagerdutyBusinessURL = compattrs.PagerdutyBusinessURL
	obj.PagerdutyURL = compattrs.PagerdutyURL
	obj.Repository = compattrs.Repository
	obj.ServiceOwner.Key = compattrs.ServiceOwner.Key
	obj.ServiceOwner.UnmarshalNFT(cid2json)
	obj.SlackChannel = compattrs.SlackChannel
}
