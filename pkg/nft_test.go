package ortelius

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNFT(t *testing.T) {

	jsonObj := []byte("Hello World")

	expected := "bafkreiffsgtnic7uebaeuaixgph3pmmq2ywglpylzwrswv5so7m23hyuny"

	var nft NFT // define user object to marshal into
	nft.Init(jsonObj)

	// fmt.Printf("%s", nft.Key)

	assert.Equal(t, expected, nft.Key, "check unmarshalled against expected results")

}
