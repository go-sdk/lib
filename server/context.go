package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/go-sdk/lib/val"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
	p map[string]string

	mu   sync.RWMutex
	keys map[string]interface{}
}

func WithContext(w http.ResponseWriter, r *http.Request, p map[string]string) *Context {
	c := &Context{W: w, R: r, p: p}
	c.keys = map[string]interface{}{}
	return c
}

func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	c.keys[key] = value
	c.mu.Unlock()
}

func (c *Context) Get(key string) (val.Value, bool) {
	c.mu.RLock()
	value, exists := c.keys[key]
	c.mu.RUnlock()
	return val.New(value), exists
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *Context) Done() <-chan struct{} {
	return nil
}

func (c *Context) Err() error {
	return nil
}

func (c *Context) Value(key interface{}) interface{} {
	if key == 0 {
		return c.R
	}
	if s, ok := key.(string); ok {
		value, _ := c.Get(s)
		return value
	}
	return nil
}
