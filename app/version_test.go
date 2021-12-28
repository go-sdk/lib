package app

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log(VERSION)
	t.Log(GITHASH)
	t.Log(BUILT)
}

func TestVersionInfo(t *testing.T) {
	t.Log(VersionInfo())
}

func TestVersionInfoMap(t *testing.T) {
	t.Log(VersionInfoMap())
}
