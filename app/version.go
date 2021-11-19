package app

import (
	"fmt"
	"runtime"
)

var (
	VERSION = ""
	GITHASH = ""
	BUILT   = ""

	GOVERSION = runtime.Version()
	GOOS      = runtime.GOOS
	GOARCH    = runtime.GOARCH
)

func VersionInfo() string {
	return fmt.Sprintf("version: %s-%s, built: %s, go: %s, os: %s, arch: %s", VERSION, GITHASH, BUILT, GOVERSION, GOOS, GOARCH)
}
