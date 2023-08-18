package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestSwagger(t *testing.T) {

	jsonObj := []byte(`{

		"content": {"openapi":"3.0.2"}
	  }`)

	expected := `{"content":{"openapi":"3.0.2"},"objtype":"Swagger"}`
	expectedCid := "bafkreigkrj4ipthicjkveygdejdt3m24vkqtz22v6ddrhoi5xtszmhk7ji"

	// define user object to marshal into
	var obj Swagger

	// convert json string into the user object
	json.Unmarshal(jsonObj, &obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(obj)
	// 	fmt.Println(cid)
	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")

}
