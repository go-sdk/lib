package dingtalk

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/httpx"
)

var (
	token  = conf.Get("dingtalk.token").String()
	secret = conf.Get("dingtalk.secret").String()
)

func TestSendMessage(t *testing.T) {
	assert.NoError(t, SendMessage(httpx.Test(), token, secret, "测试消息"))
}
