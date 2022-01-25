package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	paths = []string{
		"../internal/testdata/config.yaml",
		"../internal/testdata/config.json",
	}

	errPaths = []string{
		"../internal/testdata/config_err.json",
	}

	testPaths = []string{
		"../internal/testdata/config_clash.yaml",
	}
)

func TestLoad(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		assert.NoError(t, Load(paths...))
	})

	t.Run("with debug", func(t *testing.T) {
		assert.NoError(t, New(WithDebug(true)).Load(paths...))
	})

	t.Run("error", func(t *testing.T) {
		assert.Error(t, New(WithDebug(true)).Load(errPaths...))
	})

	t.Run("error with skip", func(t *testing.T) {
		assert.NoError(t, New(WithSkipError(true)).Load(errPaths...))
	})
}

func TestGet(t *testing.T) {
	assert.NoError(t, New(WithDebug(true), WithOverwrite(true)).Load(testPaths...))
	assert.Equal(t, int64(7890), Get("port").Int64())
	assert.Equal(t, true, Get("allow-lan").Bool())
	assert.Equal(t, false, Get("profile.store-selected").Bool())
	assert.Equal(t, []string{"114.114.114.114", "8.8.8.8"}, Get("dns.default-nameserver").SliceString())
	assert.Equal(t, []string{"obfs", "v2ray-plugin"}, Get("proxies.plugin").SliceString())
}
