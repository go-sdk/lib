package faker

import (
	"github.com/brianvoe/gofakeit/v6"
)

var F *gofakeit.Faker

func init() {
	F = gofakeit.New(0)
	gofakeit.SetGlobalFaker(F)
}
