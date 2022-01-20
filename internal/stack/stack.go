package stack

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	size = 64 << 10
	skip = 1
)

func Stack(skips ...int) string {
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	str := string(buf)
	ss := strings.Split(str, "\n")
	if len(ss) <= 1 {
		return str
	}
	res := strings.Join(ss[1+2*x(skips):], "\n")
	return ss[0] + "\n" + res
}

func FileLine(skips ...int) string {
	_, f, l, _ := runtime.Caller(x(skips))
	return fmt.Sprintf("%s:%d", f, l)
}

func FileOrLine(skips ...int) (string, int) {
	_, f, l, _ := runtime.Caller(x(skips))
	return f, l
}

func FuncName(skips ...int) string {
	pc, _, _, ok := runtime.Caller(x(skips))
	if !ok {
		return "UNKNOWN"
	}
	return runtime.FuncForPC(pc).Name()
}

func x(skips []int) int {
	if len(skips) == 0 || skips[0] < 0 {
		skips = []int{skip}
	}
	return skips[0]
}
