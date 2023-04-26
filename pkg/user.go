// Package ortelius - User defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package ortelius

import (
	"encoding/json"
)

// User defines a user
type User struct {
	Key      string `json:"_key,omitempty"`
	NftJSON  string `json:"_json,omitempty"`
	Name     string `json:"name"`
	Domain   Domain `json:"domain"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Realname string `json:"realname,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *User) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Domain   NFT    `json:"domain"`
		Email    string `json:"email,omitempty"`
		Name     string `json:"name"`
		ObjType  string `json:"objtype"`
		Phone    string `json:"phone,omitempty"`
		Realname string `json:"realname,omitempty"`
	}{
		Domain:   new(NFT).Init(obj.Domain.MarshalNFT(cid2json)), // Convert the domain object to cid+json and assign its cid to the Domain
		Email:    obj.Email,                                      // Copy
		Name:     obj.Name,                                       // Copy
		ObjType:  "User",                                         // Set
		Phone:    obj.Phone,                                      // Copy
		Realname: obj.Realname,                                   // Copy
	})

	obj.NftJSON = string(data)        // Save the json
	obj.Key = new(NFT).Init(data).Key // Calculate and save the cid for the json
	cid2json[obj.Key] = obj.NftJSON   // Add cid=json for persisting later

	return data // Return NFT Json
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *User) UnmarshalNFT(cid2json map[string]string) {
	var user User
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	err := json.Unmarshal([]byte(obj.NftJSON), &user)

	if err == nil {
		// Deep Copy
		obj.Email = user.Email
		obj.Name = user.Name
		obj.Phone = user.Phone
		obj.Realname = user.Realname
		obj.Domain.Key = user.Domain.Key
		obj.Domain.UnmarshalNFT(cid2json)
	}
}
