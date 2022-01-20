package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	t.Log(Stack())
}

func TestStack2(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(Stack(3))
		}
	}()

	panic("xxx")
}

func TestFileLine(t *testing.T) {
	t.Log(FileLine())
}

func TestFileOrLine(t *testing.T) {
	t.Log(FileOrLine())
}

func TestFuncName(t *testing.T) {
	t.Log(FuncName())
}
