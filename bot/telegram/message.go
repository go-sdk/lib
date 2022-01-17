package telegram

import (
	"fmt"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/httpx"
)

// API Document: https://core.telegram.org/bots/api#sendmessage

const Addr = "https://api.telegram.org"

const (
	ModeMarkdown   = "Markdown"
	ModeMarkdownV2 = "MarkdownV2"
	ModeHTML       = "HTML"
)

type MessageReq struct {
	ChatId                   string `json:"chat_id"`
	Text                     string `json:"text"`
	ParseMode                string `json:"parse_mode,omitempty"`
	DisableWebPagePreview    bool   `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool   `json:"disable_notification,omitempty"`
	ProtectContent           bool   `json:"protect_content,omitempty"`
	ReplyToMessageId         int64  `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply,omitempty"`
}

type MessageResp struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int64  `json:"error_code"`
	Description string `json:"description"`
}

func SendMessage(c *httpx.Client, token, to, text string) error {
	return SendMessageWithReq(c, token, &MessageReq{ChatId: to, Text: text})
}

func SendMessageWithReq(c *httpx.Client, token string, req *MessageReq) error {
	if req.Text == "" {
		req.Text = "测试消息"
	}
	resp, err := c.
		NewRequest().
		SetHeader(consts.ContentType, consts.ContentTypeJSON).
		SetBody(req).
		SetResult(&MessageResp{}).
		SetError(&MessageResp{}).
		Post(Addr + "/bot" + token + "/sendMessage")
	if err != nil {
		return fmt.Errorf("telegram: fail to call api, %s", err)
	}
	if e, ok := resp.Error().(*MessageResp); ok && !e.Ok {
		return fmt.Errorf("telegram: %d, %s", e.ErrorCode, e.Description)
	}
	return nil
}
