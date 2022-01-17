package json

import (
	"fmt"
)

const (
	prefix      = ""
	indentSpace = "  "
	indentTab   = "\t"
)

var (
	indentSpaceMap = map[int]string{
		1: " ",
		2: "  ",
		3: "   ",
		4: "    ",
	}
)

func MustMarshal(v interface{}) []byte {
	bs, _ := Marshal(v)
	return bs
}

func MustMarshalToString(v interface{}) string {
	return string(MustMarshal(v))
}

func UnmarshalFromString(s string, v interface{}) error {
	return Unmarshal([]byte(s), v)
}

func UnmarshalToMap(s string) map[string]interface{} {
	var v map[string]interface{}
	_ = UnmarshalFromString(s, &v)
	return v
}

func PrettyS(v interface{}, l ...int) string {
	indent := indentSpace
	if len(l) > 0 && l[0] >= 1 && l[0] <= 4 {
		indent = indentSpaceMap[l[0]]
	}
	bs, _ := MarshalIndent(v, prefix, indent)
	return string(bs)
}

func PrettyT(v interface{}) string {
	bs, _ := MarshalIndent(v, prefix, indentTab)
	return string(bs)
}

func Print(v interface{}) {
	fmt.Println(MustMarshalToString(v))
}

func PrintPretty(v interface{}) {
	fmt.Println(PrettyS(v))
}
