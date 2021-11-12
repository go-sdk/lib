package cron

import (
	"sync"
	"testing"
)

func Test(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)

	c := Default(nil)
	c.Add("* * * * * *", "", func() { wg.Add(-1) })
	c.Start()

	wg.Wait()

	c.Stop()
}
