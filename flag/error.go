package flag

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func IsErrFlagNotDefined(err error) bool {
	return strings.Contains(err.Error(), "flag provided but not defined")
}

func (f *FlagSet) ErrAndExit(err error, code int) {
	f.MsgAndExit(err.Error(), code)
}

func (f *FlagSet) MsgAndExit(msg string, code int) {
	if buf, ok := f.output.(*bytes.Buffer); ok {
		buf.Reset()
		fmt.Fprintln(buf, msg)
		f.usage()
		os.Stderr.Write(buf.Bytes())
	}
	os.Exit(code)
}
