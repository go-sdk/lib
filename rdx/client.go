package rdx

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"

	"github.com/go-sdk/lib/conf"
)

type Client struct {
	*redis.Client
}

var (
	mu  sync.RWMutex
	cli *Client
)

func init() {
	dsn := conf.Get("redis.dsn").String()
	if dsn == "" {
		return
	}
	x, err := New(dsn)
	if err != nil {
		panic(err)
	}
	SetDefaultCli(x)
}

func New(dsn string, def ...bool) (*Client, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	c := redis.NewClient(opt)

	err = c.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	x := &Client{Client: c}

	if len(def) == 0 || !def[0] {
		return x, nil
	}

	SetDefaultCli(x)

	return cli, nil
}

func Default() *Client {
	mu.RLock()
	defer mu.RUnlock()
	return cli
}

func SetDefaultCli(x *Client) {
	mu.Lock()
	defer mu.Unlock()
	cli = x
}
