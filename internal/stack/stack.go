package stack

import (
	"runtime"
	"strings"
)

const (
	size = 64 << 10
	skip = 3
)

func Stack() string {
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	str := string(buf)
	ss := strings.Split(str, "\n")
	if len(ss) <= 1 {
		return str
	}
	res := strings.Join(ss[1+2*skip:], "\n")
	return ss[0] + "\n" + res
}
