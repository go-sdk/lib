package json

import (
	"encoding/json"
)

var (
	Marshal       = json.Marshal
	MarshalIndent = json.MarshalIndent
	Unmarshal     = json.Unmarshal

	NewDecoder = json.NewDecoder
	NewEncoder = json.NewEncoder
)
