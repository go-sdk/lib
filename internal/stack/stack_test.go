package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(Stack())
		}
	}()

	panic("xxx")
}
