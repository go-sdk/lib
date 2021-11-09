package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.NoError(t, Load(paths...))
	})

	t.Run("with debug", func(t *testing.T) {
		assert.NoError(t, New(WithDebug()).Load(paths...))
	})

	t.Run("error", func(t *testing.T) {
		assert.Error(t, New(WithDebug()).Load(errPaths...))
	})

	t.Run("error with skip", func(t *testing.T) {
		assert.NoError(t, New(WithSkipError()).Load(errPaths...))
	})
}

func TestGet(t *testing.T) {
	assert.NoError(t, New(WithDebug(), WithOverwrite()).Load(testPaths...))
	assert.Equal(t, int64(7890), Get("port").Int64())
	assert.Equal(t, true, Get("allow-lan").Bool())
	assert.Equal(t, false, Get("profile.store-selected").Bool())
	assert.Equal(t, []string{"114.114.114.114", "8.8.8.8"}, Get("dns.default-nameserver").SliceString())
	assert.Equal(t, []string{"obfs", "v2ray-plugin"}, Get("proxies.plugin").SliceString())
}
