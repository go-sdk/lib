package srv

import (
	"github.com/gin-contrib/gzip"
)

var (
	WithExcludedExtensions  = gzip.WithExcludedExtensions
	WithExcludedPaths       = gzip.WithExcludedPaths
	WithExcludedPathsRegexs = gzip.WithExcludedPathsRegexs
	WithDecompressFn        = gzip.WithDecompressFn
)

func GZIP(options ...gzip.Option) HandlerFunc {
	return GZIPWithLevel(gzip.DefaultCompression, options...)
}

func GZIPWithLevel(level int, options ...gzip.Option) HandlerFunc {
	return gzip.Gzip(level, options...)
}
