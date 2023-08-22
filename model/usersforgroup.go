// Package model - UsersForGroup defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

// UsersForGroup defines a list of user that belong to the group
type UsersForGroup struct {
	Key      string   `json:"_key,omitempty"`
	ObjType  string   `json:"objtype,omitempty"`
	GroupKey string   `json:"group"`
	UserKeys []string `json:"users"`
}

func NewUsersForGroup() *UsersForGroup {
	return &UsersForGroup{ObjType: "UsersForGroup"}
}
