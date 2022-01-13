package testx

import (
	"testing"

	"github.com/go-sdk/lib/codec/json"
)

func Error(t *testing.T, err error, v ...interface{}) {
	t.Helper()

	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < len(v); i++ {
		t.Log(json.MustMarshalToString(v[i]))
	}
}

func Fatal(t *testing.T, err error, v ...interface{}) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
		return
	}

	for i := 0; i < len(v); i++ {
		t.Log(json.MustMarshalToString(v[i]))
	}
}
