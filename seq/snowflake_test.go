package seq

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Log(NewSnowflakeID().String())
	t.Log(NewSnowflakeID().String())
	t.Log(NewSnowflakeID().String())
}

func TestNewSnowflakeIDWithTime(t *testing.T) {
	x := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	t.Log(NewSnowflakeIDWithTime(x))
	t.Log(NewSnowflakeIDWithTime(x).Time().Local())
}

func TestNewSnowflakeIDFromString(t *testing.T) {
	s := "1212040716089630720"
	id, err := NewSnowflakeIDFromString(s)
	assert.NoError(t, err)
	assert.Equal(t, s, id.String())
}

func TestTimeFromSnowflakeIDString(t *testing.T) {
	x := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	assert.Equal(t, x.UTC(), TimeFromSnowflakeIDString(NewSnowflakeIDWithTime(x).String()).UTC())
}

func BenchmarkNewSnowflakeID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSnowflakeID()
	}
}
