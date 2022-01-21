package faker

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestF(t *testing.T) {
	t.Log(F.Name())
	t.Log(F.Email())
	t.Log(F.Emoji())
	t.Log(F.Password(true, true, true, true, false, 16))

	type User struct {
		Username string `fake:"{firstname}"`
		Password string `fake:"{password}"`
		Email    string `fake:"{email}"`
		Extra    string
	}

	user := &User{}
	require.NoError(t, F.Struct(user))
	t.Logf("%#v", user)
}
