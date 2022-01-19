package validator

import (
	"testing"
)

func Test(t *testing.T) {
	type User struct {
		Username string `validate:"required,gte=6,lte=16"`
		Password string `validate:"required,alphanum"`
	}

	user := &User{
		Username: "admin",
		Password: "p1s$w0rd",
	}

	err := V.Struct(user)

	errs := ToErrors(err)
	for i := 0; i < len(errs); i++ {
		t.Log(errs[i])
	}

	t.Log(errs.First())

	t.Log(ToErrors(err, EN).String())
}
