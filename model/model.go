// Package model - License defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import (
	"reflect"
	"strings"
)

func GetObjType(obj any) string {
	objtype := reflect.TypeOf(obj).String()

	if strings.Count(objtype, ".") > 0 {
		parts := strings.Split(objtype, ".")
		objtype = parts[len(parts)-1]
	}
	return objtype
}
