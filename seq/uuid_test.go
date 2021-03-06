package seq

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	t.Log(NewUUID().String())
	t.Log(NewUUID().String())
	t.Log(NewUUID().String())
}

func TestNewUUIDWithTime(t *testing.T) {
	x := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	t.Log(NewUUIDWithTime(x))
	t.Log(TimeFromUUIDString(NewUUIDWithTime(x).String()))
}

func TestNewUUIDFromString(t *testing.T) {
	s := "99248000-2be6-01ea-b46b-894ce5a0e50b"
	id, err := NewUUIDFromString(s)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, s, id.String())
}

func TestTimeFromUUIDString(t *testing.T) {
	x := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	assert.Equal(t, x.UTC(), TimeFromUUIDString(NewUUIDWithTime(x).String()).UTC())
}

func BenchmarkNewUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUUID()
	}
}
