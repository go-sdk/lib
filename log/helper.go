package log

import (
	"github.com/go-sdk/lib/val"
)

func ToFields(kv ...interface{}) Fields {
	if len(kv) == 0 {
		return Fields{}
	}
	if len(kv)%2 != 0 {
		kv = append(kv, "")
	}
	f := Fields{}
	for i := 0; i < len(kv); i += 2 {
		f[val.New(kv[i]).String()] = kv[i+1]
	}
	return f
}
