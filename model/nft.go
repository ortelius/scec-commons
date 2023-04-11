package model

//lint:file-ignore S1034 Ignore all assignments for switch statements

import (
	cid "github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
)

type NFT struct {
	Key string `json:"_key,omitempty"`
}

func (obj *NFT) Init(Json []byte) NFT {

	var pref = cid.Prefix{
		Version:  1,
		Codec:    uint64(mc.Raw),
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}

	_cid, err := pref.Sum(Json)

	if err != nil {
		obj.Key = ""
		return *obj
	}

	obj.Key = _cid.String()
	return *obj
}
