package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/go-sdk/lib/codec/json"
	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/httpx"
)

// API Document: https://open.dingtalk.com/document/group/custom-robot-access

const Addr = "https://oapi.dingtalk.com/robot/send"

const (
	MsgTypeText     = "text"
	MsgTypeMarkdown = "markdown"
)

type MessageReq struct {
	MsgType string     `json:"msgtype"`
	At      *MessageAt `json:"at,omitempty"`

	Text     *MessageText     `json:"text,omitempty"`
	Markdown *MessageMarkdown `json:"markdown,omitempty"`
}

type MessageAt struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	AtUserIds []string `json:"atUserIds,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type MessageText struct {
	Content string `json:"content"`
}

type MessageMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type MessageResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func SendMessage(c *httpx.Client, token, secret, text string) error {
	return SendMessageWithReq(c, token, secret, &MessageReq{MsgType: MsgTypeText, Text: &MessageText{Content: text}})
}

func SendMessageWithReq(c *httpx.Client, token, secret string, req *MessageReq) error {
	if req.MsgType == "" {
		req.MsgType = MsgTypeText
		req.Text = &MessageText{Content: "测试消息"}
	}
	return SendMessageWithMap(c, token, secret, json.UnmarshalToMap(json.MustMarshalToString(req)))
}

func SendMessageWithMap(c *httpx.Client, token, secret string, data map[string]interface{}) error {
	addr := Addr + "?access_token=" + token
	if secret != "" {
		ts, sign := signer(secret)
		addr += "&timestamp=" + ts + "&sign=" + sign
	}
	resp, err := c.
		NewRequest().
		SetHeader(consts.ContentType, consts.ContentTypeJSON).
		SetBody(data).
		SetResult(&MessageResp{}).
		Post(addr)
	if err != nil {
		return fmt.Errorf("dingtalk: fail to call api, %s", err)
	}
	if e, ok := resp.Result().(*MessageResp); ok && e.ErrCode != 0 {
		return fmt.Errorf("dingtalk: %d, %s", e.ErrCode, e.ErrMsg)
	}
	return nil
}

func signer(secret string) (string, string) {
	milli := strconv.FormatInt(time.Now().UnixMilli(), 10)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(milli + "\n" + secret))
	return milli, base64.StdEncoding.EncodeToString(h.Sum(nil))
}
