package conf

import (
	"testing"

	"github.com/go-sdk/lib/testx"
)

var (
	paths = []string{
		"../testdata/config.yaml",
		"../testdata/config.json",
	}

	errPaths = []string{
		"../testdata/config_err.json",
	}

	testPaths = []string{
		"../testdata/config_clash.yaml",
	}
)

func TestLoad(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		testx.AssertNoError(t, Load(paths...))
	})

	t.Run("with debug", func(t *testing.T) {
		testx.AssertNoError(t, New(WithDebug(true)).Load(paths...))
	})

	t.Run("error", func(t *testing.T) {
		testx.AssertError(t, New(WithDebug(true)).Load(errPaths...))
	})

	t.Run("error with skip", func(t *testing.T) {
		testx.AssertNoError(t, New(WithSkipError(true)).Load(errPaths...))
	})
}

func TestGet(t *testing.T) {
	testx.AssertNoError(t, New(WithDebug(true), WithOverwrite(true)).Load(testPaths...))
	testx.AssertEqual(t, int64(7890), Get("port").Int64())
	testx.AssertEqual(t, true, Get("allow-lan").Bool())
	testx.AssertEqual(t, false, Get("profile.store-selected").Bool())
	testx.AssertEqual(t, []string{"114.114.114.114", "8.8.8.8"}, Get("dns.default-nameserver").SliceString())
	testx.AssertEqual(t, []string{"obfs", "v2ray-plugin"}, Get("proxies.plugin").SliceString())
}
