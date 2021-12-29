package srv

import (
	"net/http"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/errx"
)

func Auth(skipPaths ...string) HandlerFunc {
	skip := map[string]struct{}{}
	for _, path := range skipPaths {
		skip[path] = struct{}{}
	}

	return func(c *Context) {
		if _, ok := skip[c.FullPath()]; ok {
			c.Next()
			return
		}

		if _, ok := c.Get(consts.CTokenRaw); !ok {
			c.JSON(http.StatusOK, errx.Unauthorized("missing "+consts.Authorization))
			c.Abort()
			return
		}

		if _, ok := c.Get(consts.CToken); !ok {
			c.JSON(http.StatusOK, errx.Unauthorized("invalid "+consts.Authorization))
			c.Abort()
			return
		}
	}
}
