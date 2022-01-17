package telegram

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/httpx"
)

var (
	token = conf.Get("telegram.token").String()
	to    = conf.Get("telegram.to").String()
)

func TestSendMessage(t *testing.T) {
	assert.NoError(t, SendMessage(httpx.Test(), token, to, "测试消息"))
}
