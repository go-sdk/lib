//go:build jsoniter
// +build jsoniter

package json

import (
	"github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	Marshal       = json.Marshal
	MarshalIndent = json.MarshalIndent
	Unmarshal     = json.Unmarshal

	NewDecoder = json.NewDecoder
	NewEncoder = json.NewEncoder
)
