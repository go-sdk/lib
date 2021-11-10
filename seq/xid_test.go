package seq

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewXID(t *testing.T) {
	t.Log(NewXID().String())
	t.Log(NewXID().String())
	t.Log(NewXID().String())
}

func TestNewXIDWithTime(t *testing.T) {
	x := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	t.Log(NewXIDWithTime(x))
	t.Log(NewXIDWithTime(x).Time().Local())
}

func TestNewXIDFromString(t *testing.T) {
	s := "71md606labs31jdhtarg"
	id, err := NewXIDFromString(s)
	assert.NoError(t, err)
	assert.Equal(t, s, id.String())
}

func TestTimeFromXIDString(t *testing.T) {
	x := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	assert.Equal(t, x.UTC(), TimeFromXIDString(NewXIDWithTime(x).String()).UTC())
}

func BenchmarkNewXID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewXID()
	}
}
