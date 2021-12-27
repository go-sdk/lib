package db

import (
	"context"

	"github.com/go-sdk/lib/seq"
)

type idGeneratorFunc func(ctx context.Context) string

var idGenerator idGeneratorFunc = func(ctx context.Context) string { return seq.NewSnowflakeID().String() }

func SetIDGenerator(f idGeneratorFunc) { idGenerator = f }
