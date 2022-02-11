package token

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/crypto"
)

var (
	key    = []byte(conf.Get("token.key").StringD("99248000-2be6-01ea-b46b-894ce5a0e50b"))
	expire = conf.Get("token.expire").DurationD(7 * 24 * time.Hour)

	bearerPrefix    = consts.Bearer + " "
	bearerPrefixLen = len(bearerPrefix)

	mu = sync.Mutex{}
)

func GetExpire() time.Duration {
	mu.Lock()
	defer mu.Unlock()
	return expire
}

func SetExpire(d time.Duration) {
	mu.Lock()
	defer mu.Unlock()
	expire = d
}

type Token struct {
	c *claim

	raw string
}

type Extra map[string]interface{}

type claim struct {
	ID       string `json:"jti,omitempty"`
	Issuer   string `json:"iss,omitempty"`
	UserId   string `json:"uid,omitempty"`
	IssuedAt int64  `json:"iat,omitempty"`

	Extra `json:"exa,omitempty"`
}

func (c *claim) Valid() error {
	return nil
}

type Option func(t *Token)

func WithID(id string) Option {
	return func(t *Token) {
		t.c.ID = id
	}
}

func WithExtra(extra Extra) Option {
	return func(t *Token) {
		for k, v := range extra {
			t.c.Extra[k] = v
		}
	}
}

func New(iss, uid string, iat int64, opts ...Option) *Token {
	id := crypto.RandString(8, crypto.CharsetLetterLower)

	if iat == 0 {
		iat = time.Now().Unix()
	}

	token := &Token{
		c: &claim{
			ID:       id,
			Issuer:   iss,
			UserId:   uid,
			IssuedAt: iat,
			Extra:    Extra{},
		},
	}

	for i := 0; i < len(opts); i++ {
		opts[i](token)
	}

	return token
}

func (t *Token) Sign() (string, error) {
	c := jwt.NewWithClaims(jwt.SigningMethodHS256, t.c)
	return c.SignedString(key)
}

func (t *Token) SignString() string {
	s, _ := t.Sign()
	return s
}

type tokenKey struct{}

func (t *Token) WithContext(ctx ...context.Context) context.Context {
	if len(ctx) == 0 || ctx[0] == nil {
		ctx = []context.Context{context.Background()}
	}
	return context.WithValue(ctx[0], tokenKey{}, t)
}

func (t *Token) GetID() string {
	return t.c.ID
}

func (t *Token) GetIssuer() string {
	return t.c.Issuer
}

func (t *Token) GetUserId() string {
	return t.c.UserId
}

func (t *Token) GetIssuedAt() int64 {
	return t.c.IssuedAt
}

func (t *Token) GetExtra() Extra {
	return t.c.Extra
}

func (t *Token) IsExpired() bool {
	if t.GetIssuedAt() == -1 {
		return false
	}
	return time.Now().Add(-GetExpire()).Unix() > t.GetIssuedAt()
}

func (t *Token) Refresh() *Token {
	if t.GetIssuedAt() == -1 {
		return t
	}
	return New(t.GetIssuer(), t.GetUserId(), time.Now().Unix(), WithID(t.GetID()), WithExtra(t.GetExtra()))
}

func Parse(s string) (*Token, error) {
	if strings.HasPrefix(strings.ToLower(s), bearerPrefix) {
		s = s[bearerPrefixLen:]
	}
	t := &Token{c: &claim{}, raw: s}
	_, err := jwt.ParseWithClaims(s, t.c, func(t *jwt.Token) (interface{}, error) { return key, nil })
	return t, err
}

var ErrNotParsed = fmt.Errorf("token: not parsed")

func FromContext(ctx context.Context) (*Token, error) {
	if t, ok := ctx.Value(tokenKey{}).(*Token); ok && t.c != nil {
		return t, nil
	}
	return nil, ErrNotParsed
}

func MustFromContext(ctx context.Context) *Token {
	t, err := FromContext(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

func GetID(ctx context.Context) string {
	return MustFromContext(ctx).GetID()
}

func GetIssuer(ctx context.Context) string {
	return MustFromContext(ctx).GetIssuer()
}

func GetUserId(ctx context.Context) string {
	return MustFromContext(ctx).GetUserId()
}

func GetIssuedAt(ctx context.Context) int64 {
	return MustFromContext(ctx).GetIssuedAt()
}

func GetExtra(ctx context.Context) Extra {
	return MustFromContext(ctx).GetExtra()
}
