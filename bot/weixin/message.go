package weixin

import (
	"fmt"

	"github.com/go-sdk/lib/codec/json"
	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/httpx"
)

// API Document: https://developer.work.weixin.qq.com/document/path/90236

const Addr = "https://qyapi.weixin.qq.com/cgi-bin"

const (
	MsgTypeText     = "text"
	MsgTypeMarkdown = "markdown"
)

type MessageReq struct {
	ToUser                 string `json:"touser,omitempty"`
	ToParty                string `json:"toparty,omitempty"`
	ToTag                  string `json:"totag,omitempty"`
	MsgType                string `json:"msgtype"`
	AgentId                int64  `json:"agentid"`
	Safe                   int64  `json:"safe,omitempty"`
	EnableIdTrans          int64  `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int64  `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int64  `json:"duplicate_check_interval,omitempty"`

	Text     *MessageText     `json:"text,omitempty"`
	Markdown *MessageMarkdown `json:"markdown,omitempty"`
}

type MessageText struct {
	Content string `json:"content"`
}

type MessageMarkdown struct {
	Content string `json:"content"`
}

type MessageResp struct {
	ErrCode      int64  `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
	MsgId        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

func SendMessage(c *httpx.Client, aid int64, token, to string, text string) error {
	return SendMessageWithReq(c, token, &MessageReq{ToUser: to, MsgType: MsgTypeText, AgentId: aid, Text: &MessageText{Content: text}})
}

func SendMessageWithReq(c *httpx.Client, token string, req *MessageReq) error {
	if req.MsgType == "" {
		req.MsgType = MsgTypeText
		req.Text = &MessageText{Content: "测试消息"}
	}
	return SendMessageWithMap(c, token, json.UnmarshalToMap(json.MustMarshalToString(req)))
}

func SendMessageWithMap(c *httpx.Client, token string, data map[string]interface{}) error {
	resp, err := c.
		NewRequest().
		SetHeader(consts.ContentType, consts.ContentTypeJSON).
		SetBody(data).
		SetResult(&MessageResp{}).
		Post(Addr + "/message/send?access_token=" + token)
	if err != nil {
		return fmt.Errorf("weixin: fail to call api, %s", err)
	}
	if e, ok := resp.Result().(*MessageResp); ok && e.ErrCode != 0 {
		return fmt.Errorf("weixin: %d, %s", e.ErrCode, e.ErrMsg)
	}
	return nil
}

type TokenResp struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetToken(c *httpx.Client, id, secret string) (*TokenResp, error) {
	resp, err := c.
		NewRequest().
		SetResult(&TokenResp{}).
		Get(Addr + "/gettoken?corpid=" + id + "&corpsecret=" + secret)
	if err != nil {
		return nil, fmt.Errorf("weixin: fail to call api, %s", err)
	}
	return resp.Result().(*TokenResp), nil
}
