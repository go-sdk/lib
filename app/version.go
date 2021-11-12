package app

import (
	"fmt"
)

var (
	VERSION = ""
	GITHASH = ""
	BUILT   = ""
)

func VersionInfo() string {
	return fmt.Sprintf("version: %s-%s, built: %s", VERSION, GITHASH, BUILT)
}
