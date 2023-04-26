package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuditLog(t *testing.T) {
	cid2json := make(map[string]string, 0)

	jsonObj := []byte(`{
		"_key": "bafkreifkib6s6xfvscqpeby574sikl4d43w6bz5naqbpb4htqr3s4sdbk4",
		"auditlog": [{
				"_key": "bafkreibpkskhzdvulykusspc7xfvpd3tflmiz6yftdhdi2a7d6xtipqbqa",
				"action": "Created",
				"user": {
					"_key": "bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi",
					"domain": {
						"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
						"name": "GLOBAL"
					},
					"email": "admin@ortelius.io",
					"name": "admin",
					"phone": "505-444-5566",
					"realname": "Ortelius Admin"
				},
				"when": "2023-04-23T10:20:30.400+02:30"
			},
			{
				"_key": "bafkreibpkskhzdvulykusspc7xfvpd3tflmiz6yftdhdi2a7d6xtipqbqa",
				"action": "Updated",
				"user": {
					"_key": "bafkreiaj3gyc7k2gqs7roc6rduasmt4htgjagrqfulo2cd566xk3tei6zi",
					"domain": {
						"_key": "bafkreicjtrtqndgtn37wc2up26sombgyh6uqwnn4orarfdqyw63lvg5aty",
						"name": "GLOBAL"
					},
					"email": "admin@ortelius.io",
					"name": "admin",
					"phone": "505-444-5566",
					"realname": "Ortelius Admin"
				},
				"when": "2023-05-23T10:20:30.400+02:30"
			}
		]
	}`)

	expected := `{"auditlog":[{"_key":"bafkreibpkskhzdvulykusspc7xfvpd3tflmiz6yftdhdi2a7d6xtipqbqa"},{"_key":"bafkreieczryfb3ab4manho4jnm3sj7kpkojmlte2h3pkshc7qt4sqw5354"}]}`

	var audit2nft AuditLog // define user object to marshal into

	json.Unmarshal(jsonObj, &audit2nft) // convert json string into the user object
	audit2nft.MarshalNFT(cid2json)      // generate the cid and nft json for user object
	// fmt.Printf("%s=%s\n", audit2nft.Key, audit2nft.NftJSON)
	assert.Equal(t, audit2nft.NftJSON, expected, "check nft json against expected results")

	var nft2audit AuditLog // define user object to marshal into

	nft2audit.NftJSON = expected     // set the nft json
	nft2audit.UnmarshalNFT(cid2json) // convert the json string into the user object
	nft2audit.MarshalNFT(cid2json)   // recalcuate the cid and nft json for the new user object
	assert.Equal(t, nft2audit.NftJSON, expected, "check unmarshalled user against expected results")

}
