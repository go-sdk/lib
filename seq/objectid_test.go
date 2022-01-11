package seq

import (
	"testing"
	"time"

	"github.com/go-sdk/lib/testx"
)

func TestNewObjectID(t *testing.T) {
	t.Log(NewObjectID().String())
	t.Log(NewObjectID().String())
	t.Log(NewObjectID().String())
}

func TestNewObjectIDWithTime(t *testing.T) {
	x := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	t.Log(NewObjectIDWithTime(x))
	t.Log(NewObjectIDWithTime(x).Time().Local())
}

func TestNewObjectIDFromString(t *testing.T) {
	s := "618b8c8048e378f5036ec626"
	id, err := NewObjectIDFromString(s)
	testx.AssertNoError(t, err)
	testx.AssertEqual(t, s, id.String())
}

func TestTimeFromObjectIDString(t *testing.T) {
	x := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	testx.AssertEqual(t, x.UTC(), TimeFromObjectIDString(NewObjectIDWithTime(x).String()).UTC())
}

func BenchmarkNewObjectID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewObjectID()
	}
}
