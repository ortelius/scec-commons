// Package ortelius - NFT defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package ortelius

//lint:file-ignore S1034 Ignore all assignments for switch statements

import (
	cid "github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
)

// NFT defines the CID key for a JSON string
type NFT struct {
	Key string `json:"_key,omitempty"`
}

// Init is the constructor for the NFT struct.
// It calculates the CID for the JSON string
// and saves it as the Key.
func (obj *NFT) Init(jsonStr []byte) NFT {

	var pref = cid.Prefix{
		Version:  1,
		Codec:    uint64(mc.Raw),
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}

	_cid, err := pref.Sum(jsonStr)

	if err != nil {
		obj.Key = ""
		return *obj
	}

	obj.Key = _cid.String()
	return *obj
}
