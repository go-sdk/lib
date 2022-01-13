package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, DataRaw, string(MustMarshal(Data)))
}

func TestMustMarshalToString(t *testing.T) {
	assert.Equal(t, DataRaw, MustMarshalToString(Data))
}

func TestUnmarshalFromString(t *testing.T) {
	var v interface{}
	assert.NoError(t, UnmarshalFromString(DataRaw, &v))
	assert.Equal(t, DataRaw, MustMarshalToString(v))
}

func TestPrettyS(t *testing.T) {
	assert.Equal(t, DataPrettySRaw, PrettyS(Data))
	assert.Equal(t, DataPrettySRaw, PrettyS(Data, 2))
}

func TestPrettyT(t *testing.T) {
	assert.Equal(t, DataPrettyTRaw, PrettyT(Data))
}

func TestPrint(t *testing.T) {
	Print(Data)
}

func TestPrintPretty(t *testing.T) {
	PrintPretty(Data)
}
