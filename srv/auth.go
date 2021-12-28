package srv

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/errx"
	"github.com/go-sdk/lib/token"
)

func Auth(skipPaths ...string) HandlerFunc {
	skip := map[string]struct{}{}
	for _, path := range skipPaths {
		skip[path] = struct{}{}
	}

	return func(c *gin.Context) {
		if _, ok := skip[c.FullPath()]; ok {
			c.Next()
			return
		}

		s := c.GetString(consts.Authorization)
		if s == "" {
			s = c.GetHeader(consts.Authorization)
			if s == "" {
				s, _ = c.GetQuery(consts.Authorization)
			}
		}

		if s == "" {
			c.JSON(http.StatusOK, errx.Unauthorized("missing "+consts.Authorization))
			c.Abort()
			return
		}

		ss := strings.Split(s, " ")
		if len(ss) > 1 {
			s = ss[len(ss)-1]
		}

		t, err := token.Parse(s)
		if err != nil {
			c.JSON(http.StatusOK, errx.Unauthorized(err.Error()))
			c.Abort()
			return
		}

		c.Set(consts.CToken, t)
		c.Set(consts.CTokenRaw, s)
	}
}
