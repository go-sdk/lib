package weixin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/httpx"
)

var (
	id     = conf.Get("weixin.id").String()
	secret = conf.Get("weixin.secret").String()
	aid    = conf.Get("weixin.aid").Int64()
)

func TestSendMessage(t *testing.T) {
	resp, err := GetToken(httpx.Test(), id, secret)
	require.NoError(t, err)

	assert.NoError(t, SendMessage(httpx.Test(), aid, resp.AccessToken, "@all", "测试消息"))
}
