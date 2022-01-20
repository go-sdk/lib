package errx

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/internal/stack"
	"github.com/go-sdk/lib/log"
)

var (
	track = conf.Get("err.track").Bool()
	exit  = conf.Get("err.exit").Bool()

	mu = sync.Mutex{}

	skip = 2
	line = 3
)

func SetTrack(t bool) {
	mu.Lock()
	track = t
	mu.Unlock()
}

func SetNotEmptyExit(t bool) {
	mu.Lock()
	exit = t
	mu.Unlock()
}

func NotEmpty(err error, skips ...int) bool {
	if !track {
		return err != nil
	}

	k := x(skips)

	f, l := stack.FileOrLine(k)
	fn := stack.FuncName(k)

	sb := strings.Builder{}
	sb.WriteString("\n")
	sb.WriteString(fn)
	sb.WriteString("(): ")
	sb.WriteString("\x1b[31m")
	sb.WriteString(err.Error())
	sb.WriteString("\x1b[0m")
	sb.WriteString("\nline ")
	sb.WriteString(strconv.FormatInt(int64(l), 10))
	sb.WriteString(" of ")
	sb.WriteString(f)
	sb.WriteString("\n")

	fileBS, err := ioutil.ReadFile(f)
	if err == nil {
		lines := strings.Split(string(fileBS), "\n")
		for i := l - line; i < l+line && i < len(lines); i++ {
			if i == l-1 {
				sb.WriteString("\x1b[31m")
			}
			sb.WriteString(fmt.Sprintf("%4d", i+1))
			sb.WriteString(": ")
			sb.WriteString(lines[i])
			sb.WriteString("\n")
			if i == l-1 {
				sb.WriteString("\x1b[0m")
			}
		}
	}

	log.Error(sb.String())

	if exit {
		os.Exit(1)
	}

	return true
}

func x(skips []int) int {
	if len(skips) == 0 || skips[0] < 0 {
		skips = []int{skip}
	}
	return skips[0]
}
