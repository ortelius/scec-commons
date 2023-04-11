package model

import "encoding/json"

type CompAttrs struct {
	Key                  string `json:"_key,omitempty"`
	NftJson              string `json:"_json,omitempty"`
	BuildDate            string `json:"builddate,omitempty"`
	BuildId              string `json:"buildid,omitempty"`
	BuildUrl             string `json:"buildurl,omitempty"`
	Chart                string `json:"chart,omitempty"`
	ChartNamespace       string `json:"chartnamespace,omitempty"`
	ChartRepo            string `json:"chartrepo,omitempty"`
	ChartRepoUrl         string `json:"chartrepourl,omitempty"`
	ChartVersion         string `json:"chartversion,omitempty"`
	DiscordChannel       string `json:"discordchannel,omitempty"`
	DockerRepo           string `json:"dockerrepo,omitempty"`
	DockerSha            string `json:"dockersha,omitempty"`
	DockerTag            string `json:"dockertag,omitempty"`
	GitCommit            string `json:"gitcommit,omitempty"`
	GitRepo              string `json:"gitrepo,omitempty"`
	GitTag               string `json:"gittag,omitempty"`
	GitUrl               string `json:"giturl,omitempty"`
	HipchatChannel       string `json:"hipchatchannel,omitempty"`
	PagerdutyBusinessUrl string `json:"pagerdutybusinessurl,omitempty"`
	PagerdutyUrl         string `json:"pagerdutyurl,omitempty"`
	Repository           string `json:"repository,omitempty"`
	ServiceOwner         User   `json:"serviceowner,omitempty"`
	SlackChannel         string `json:"slackchannel,omitempty"`
}

func (obj *CompAttrs) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		BuildDate            string `json:"builddate,omitempty"`
		BuildId              string `json:"buildid,omitempty"`
		BuildUrl             string `json:"buildurl,omitempty"`
		Chart                string `json:"chart,omitempty"`
		ChartNamespace       string `json:"chartnamespace,omitempty"`
		ChartRepo            string `json:"chartrepo,omitempty"`
		ChartRepoUrl         string `json:"chartrepourl,omitempty"`
		ChartVersion         string `json:"chartversion,omitempty"`
		DiscordChannel       string `json:"discordchannel,omitempty"`
		DockerRepo           string `json:"dockerrepo,omitempty"`
		DockerSha            string `json:"dockersha,omitempty"`
		DockerTag            string `json:"dockertag,omitempty"`
		GitCommit            string `json:"gitcommit,omitempty"`
		GitRepo              string `json:"gitrepo,omitempty"`
		GitTag               string `json:"gittag,omitempty"`
		GitUrl               string `json:"giturl,omitempty"`
		HipchatChannel       string `json:"hipchatchannel,omitempty"`
		ObjType              string `json:"objtype"`
		PagerdutyBusinessUrl string `json:"pagerdutybusinessurl,omitempty"`
		PagerdutyUrl         string `json:"pagerdutyurl,omitempty"`
		Repository           string `json:"repository,omitempty"`
		ServiceOwner         NFT    `json:"serviceowner,omitempty"`
		SlackChannel         string `json:"slackchannel,omitempty"`
	}{
		BuildDate:            obj.BuildDate,
		BuildId:              obj.BuildId,
		BuildUrl:             obj.BuildUrl,
		Chart:                obj.Chart,
		ChartNamespace:       obj.ChartNamespace,
		ChartRepo:            obj.ChartRepo,
		ChartRepoUrl:         obj.ChartRepoUrl,
		ChartVersion:         obj.ChartVersion,
		DiscordChannel:       obj.DiscordChannel,
		DockerRepo:           obj.DockerRepo,
		DockerSha:            obj.DockerSha,
		DockerTag:            obj.DockerTag,
		GitCommit:            obj.GitCommit,
		GitRepo:              obj.GitRepo,
		GitTag:               obj.GitTag,
		GitUrl:               obj.GitUrl,
		HipchatChannel:       obj.HipchatChannel,
		ObjType:              "CompAttr",
		PagerdutyBusinessUrl: obj.PagerdutyBusinessUrl,
		PagerdutyUrl:         obj.PagerdutyUrl,
		Repository:           obj.Repository,
		ServiceOwner:         new(NFT).Init(obj.ServiceOwner.MarshalNFT(cid2json)),
		SlackChannel:         obj.SlackChannel,
	})

	obj.NftJson = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJson // Add cid=json for persisting later

	return data
}

func (obj *CompAttrs) UnmarshalNFT(cid2json map[string]string) {
	var compattrs CompAttrs
	var exists bool
	var NftJson string

	// get the json from storage
	if NftJson, exists = cid2json[obj.Key]; exists {
		obj.NftJson = NftJson // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJson), &compattrs)

	// Deep Copy
	obj.BuildDate = compattrs.BuildDate
	obj.BuildId = compattrs.BuildId
	obj.BuildUrl = compattrs.BuildUrl
	obj.Chart = compattrs.Chart
	obj.ChartNamespace = compattrs.ChartNamespace
	obj.ChartRepo = compattrs.ChartRepo
	obj.ChartRepoUrl = compattrs.ChartRepoUrl
	obj.ChartVersion = compattrs.ChartVersion
	obj.DiscordChannel = compattrs.DiscordChannel
	obj.DockerRepo = compattrs.DockerRepo
	obj.DockerSha = compattrs.DockerSha
	obj.DockerTag = compattrs.DockerTag
	obj.GitCommit = compattrs.GitCommit
	obj.GitRepo = compattrs.GitRepo
	obj.GitTag = compattrs.GitTag
	obj.GitUrl = compattrs.GitUrl
	obj.HipchatChannel = compattrs.HipchatChannel
	obj.PagerdutyBusinessUrl = compattrs.PagerdutyBusinessUrl
	obj.PagerdutyUrl = compattrs.PagerdutyUrl
	obj.Repository = compattrs.Repository
	obj.ServiceOwner.Key = compattrs.ServiceOwner.Key
	obj.ServiceOwner.UnmarshalNFT(cid2json)
	obj.SlackChannel = compattrs.SlackChannel
}
