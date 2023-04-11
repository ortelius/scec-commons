package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponentVersionDetails(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreib4oglwtfgurisidrzqyppgx3tl5nuxic6fl4subb2dvkalb4hyqi",
		"name": "Hello World;v1.0.0",
		"domain": {
			"_key": "bafkreih5u7cqrnv5oc2xutjhzylffaw7xvlw5nvthtlb5mg43s7wazgxle",
			"name": "GLOBAL.My Project"
		},
		"parent_key": "",
		"predecessor_key": "",
		"owner": {
			"_key": "bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi",
			"name": "admin",
			"domain": {
				"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
				"name": "GLOBAL"
			},
			"email": "admin@ortelius.io",
			"phone": "505-444-5566",
			"realname": "Ortelius Admin"
		},
		"creator": {
			"_key": "bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi",
			"name": "admin",
			"domain": {
				"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
				"name": "GLOBAL"
			},
			"email": "admin@ortelius.io",
			"phone": "505-444-5566",
			"realname": "Ortelius Admin"
		},
		"created": "2023-04-23T10:20:30.400+02:30",
		"comptype": "docker",
		"packages": {
			"_key": "bafkreia455fosaucpob7sebue7ae5pojh2gjrc5fijytq5ruunhn7ndfpi",
			"packages": [{
					"_key": "bafkreie72z3l77p6nkpkrmfyxqxopwnjiq3ztkur7ayhloleyqljsuf5ve",
					"purl": "pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
					"name": "libc-bin",
					"version": "2.19.18+deb8u7",
					"license_key": 23,
					"license": "GP-2.0"
				},
				{
					"_key": "bafkreianebpkdcvcna7ewjmpcspbw7k67lpf6oiuawdizxuwv6gnojrdla",
					"purl": "pkg:deb/debian/libcpp-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
					"name": "libcpp-bin",
					"version": "2.19.18+deb8u7",
					"license_key": 23,
					"license": "GP-2.0"
				}
			]
		},
		"vulnerabilties": {
			"_key": "bafkreigt6hy7df5mqqo7akcvkmu62aovlrv7sm7mqw6udp2f3777czh5i4",
			"vulnerabilties": [{
					"_key": "bafkreie4t6xqy65fsykdtbz5rcwjivii7rpapjq55zwnpnbl6uj6dysdbu",
					"name": "CVE-1823"
				},
				{
					"_key": "bafkreifpc2htaxkfttbyzywe4nsn2kbicq6w5rj2sqf2lr6nw5qw5qqzka",
					"name": "CVE-1824"
				}
			]
		},
		"readme": {
			"_key": "bafkreigp3a4djvcp5uidon4366bnggouiw4ghts6ngwfgt5jncbtj6lshy",
			"content": ["# README", "## Sample"]
		},
		"license": {
			"_key": "bafkreicv5dl7ozitvglmfm6jsqvw3f2pqqwpocgl6vbefgratj2gohtvmy",
			"content": ["# Apache 2", "## Summary"]
		},
		"swagger": {
			"_key": "bafkreihntxpcz47l7blasyee3tv2as6qthk44yifi7blwswi6w5dbnpcza",
			"content": ["# Rest APIs", "## GET /user"]
		},
		"applications": {
			"_key": "bafkreicecnx2gvntm6fbcrvnc336qze6st5u7qq7457igegamd3bzkx7ri"
			"applications": []
		},
		"providing": {
			"_key": "bafkreifpq5fvyuajskoik4j7n362edrr6ubkoxj5gfm74gk2lxdpsw2fmy",
			"provides": ["/user"]
		},
		"consuming": {
			"_key": "bafkreibz4duaceeggbwnl7zhvzqbttyekglpwzac4dr57ig37fzwjvdcaa",
			"consumes": ["/user"]
		},
		"attrs": {
			"_key": "bafkreihvif2vgdjswemtay7urmaehnzwuzo2qr2sgbfivk737fxdv6juda",
			"builddate": "Mon Jan 31 16:18:26 2022",
			"build_key": "178",
			"buildurl": "https://circleci.com/gh/ortelius/store-cartservice/178",
			"chart": "chart/ms-cartservice",
			"chartnamespace": "default",
			"chartrepo": "msproject/ms-chartservice",
			"chartrepourl": "https://helm.msprogject/stable/msproject/ms-chartservice",
			"chartversion": "1.0.0",
			"discordchannel": "https://discord.gg/A4hx3",
			"dockerrepo": "myproject/ms-chartservice",
			"dockersha": "5d3d677e1",
			"dockertag": "v1.0.0",
			"gitcommit": "2adc111",
			"gitrepo": "msproject/ms-chartservice",
			"gittag": "main",
			"giturl": "https://github.com/msproject/ms-chartservice",
			"hipchatchannel": "",
			"pagerdutybusinessurl": "https://pagerduty.com/business/ms-chartservice",
			"pagerdutyurl": "https://pagerduty.com/business/ms-chartservice",
			"serviceowner": {
				"_key": "bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi",
				"name": "admin",
				"domain": {
					"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
					"name": "GLOBAL"
				},
				"email": "admin@ortelius.io",
				"phone": "505-444-5566",
				"realname": "Ortelius Admin"
			},
			"slackchannel": "https://myproject.slack.com/444aaa"
		}
	}`)

	expected := `{"applications":{"_key":"bafkreicecnx2gvntm6fbcrvnc336qze6st5u7qq7457igegamd3bzkx7ri"},"attrs":{"_key":"bafkreigomg3gqp3k732jjg63bryzi5xlwjs74kw5ijte2vqmrr4qffz4ai"},"autditlog":{"_key":"bafkreicecnx2gvntm6fbcrvnc336qze6st5u7qq7457igegamd3bzkx7ri"},"comptype":"","consuming":{"_key":"bafkreigstwwmxwvxzgpbq7reb3aakqyewdr4haucw6lh6rszew7hysfl2m"},"created":"0001-01-01T00:00:00Z","creator":{"_key":"bafkreicyvutyito2jltc62vvvudogzk7zyint3dq6q32nhyackeekokiti"},"domain":{"_key":"bafkreiafns7npg5wgy7hzypm255tbfeyws552tx2sgksm2ta74o3mhlxnu"},"license":{"_key":"bafkreigrdojynv5hckqqh46hyrs6a2mfbuxdqdjn6lfbdaopjhuddhgcua"},"name":"","objtype":"ComponentVersionDetails","owner":{"_key":"bafkreicyvutyito2jltc62vvvudogzk7zyint3dq6q32nhyackeekokiti"},"packages":{"_key":"bafkreicecnx2gvntm6fbcrvnc336qze6st5u7qq7457igegamd3bzkx7ri"},"providing":{"_key":"bafkreid3dyog5ulkyb7wihve2gzgs34zij4f2v5jvup24qcoxrzcb6sqgq"},"readme":{"_key":"bafkreidp5ldgzt3yq3ktdtvktbszmhsouvizq34kzobbchhpa5ztdkp42e"},"swagger":{"_key":"bafkreifgnk7ql6jsh35wxzhxqagirugmxr22lrnv3cvpqc5sowikqiw7pa"},"vulnerabilities":{"_key":"bafkreicecnx2gvntm6fbcrvnc336qze6st5u7qq7457igegamd3bzkx7ri"}}`

	var compver2nft ComponentVersionDetails // define user object to marshal into

	json.Unmarshal(jsonObj, &compver2nft) // convert json string into the user object
	compver2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", compver2nft.Key, compver2nft.NftJson)
	assert.Equal(t, expected, compver2nft.NftJson, "check nft json against expected results")

	var nft2compver ComponentVersionDetails // define user object to marshal into

	nft2compver.NftJson = expected     // set the nft json
	nft2compver.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2compver.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, expected, nft2compver.NftJson, "check unmarshalled user against expected results")

}
