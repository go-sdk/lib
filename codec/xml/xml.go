package xml

import (
	"encoding/xml"
)

var (
	Marshal       = xml.Marshal
	MarshalIndent = xml.MarshalIndent
	Unmarshal     = xml.Unmarshal

	NewDecoder = xml.NewDecoder
	NewEncoder = xml.NewEncoder

	CopyToken       = xml.CopyToken
	NewTokenDecoder = xml.NewTokenDecoder
)
