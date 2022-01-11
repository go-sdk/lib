package json

import (
	"testing"

	"github.com/go-sdk/lib/testx"
)

var (
	Data = struct {
		A string
		B bool
		C float64
		D int64
		E map[string]interface{}
		F []interface{}
	}{
		A: "json",
		B: true,
		C: 1234.5678,
		D: 987654321,
		E: map[string]interface{}{"x": 1, "y": 2},
		F: []interface{}{1, 2, 3},
	}

	DataRaw = `{"A":"json","B":true,"C":1234.5678,"D":987654321,"E":{"x":1,"y":2},"F":[1,2,3]}`

	DataPrettySRaw = `{
  "A": "json",
  "B": true,
  "C": 1234.5678,
  "D": 987654321,
  "E": {
    "x": 1,
    "y": 2
  },
  "F": [
    1,
    2,
    3
  ]
}`

	DataPrettyTRaw = `{
	"A": "json",
	"B": true,
	"C": 1234.5678,
	"D": 987654321,
	"E": {
		"x": 1,
		"y": 2
	},
	"F": [
		1,
		2,
		3
	]
}`
)

func TestMustMarshal(t *testing.T) {
	testx.AssertEqual(t, DataRaw, string(MustMarshal(Data)))
}

func TestMustMarshalToString(t *testing.T) {
	testx.AssertEqual(t, DataRaw, MustMarshalToString(Data))
}

func TestUnmarshalFromString(t *testing.T) {
	var v interface{}
	testx.AssertNoError(t, UnmarshalFromString(DataRaw, &v))
	testx.AssertEqual(t, DataRaw, MustMarshalToString(v))
}

func TestPrettyS(t *testing.T) {
	testx.AssertEqual(t, DataPrettySRaw, PrettyS(Data))
	testx.AssertEqual(t, DataPrettySRaw, PrettyS(Data, 2))
}

func TestPrettyT(t *testing.T) {
	testx.AssertEqual(t, DataPrettyTRaw, PrettyT(Data))
}

func TestPrint(t *testing.T) {
	Print(Data)
}

func TestPrintPretty(t *testing.T) {
	PrintPretty(Data)
}
