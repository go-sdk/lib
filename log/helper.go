package log

import (
	"fmt"
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
		f[fmt.Sprintf("%v", kv[i])] = kv[i+1]
	}
	return f
}
