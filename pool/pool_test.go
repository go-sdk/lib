package pool

import (
	"context"
	"fmt"
	"testing"
	"time"
)

var sleep = 200 * time.Millisecond

func TestNew(t *testing.T) {
	p := New()
	p.Add(func(ctx context.Context) error { t.Log(1); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(2); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(3); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(4); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(5); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(6); time.Sleep(sleep); return nil })
	t.Log(p.Run(2))
}

func TestNewWithPanic(t *testing.T) {
	p := New()
	p.Add(func(ctx context.Context) error { t.Log(1); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(2); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(3); time.Sleep(sleep); panic("panic") })
	p.Add(func(ctx context.Context) error { t.Log(4); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(5); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(6); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(7); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(8); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(9); time.Sleep(sleep); return nil })
	t.Log(p.Run(2))
}

func TestNewWithError(t *testing.T) {
	p := New()
	p.Add(func(ctx context.Context) error { t.Log(1); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(2); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(3); time.Sleep(sleep); return fmt.Errorf("error") })
	p.Add(func(ctx context.Context) error { t.Log(4); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(5); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(6); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(7); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(8); time.Sleep(sleep); return nil })
	p.Add(func(ctx context.Context) error { t.Log(9); time.Sleep(sleep); return nil })
	t.Log(p.Run(2))
}
