package srv

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/consts"
	"github.com/go-sdk/lib/log"
	"github.com/go-sdk/lib/seq"
	"github.com/go-sdk/lib/token"
)

type (
	Engine  = gin.Engine
	Context = gin.Context

	RouterGroup = gin.RouterGroup

	HandlerFunc = gin.HandlerFunc
)

func init() {
	gin.SetMode(strings.ToLower(conf.Get("srv.gin.mode").StringD(gin.ReleaseMode)))
}

func New() *Engine {
	e := gin.New()
	e.Use(Init())
	return e
}

func Default() *Engine {
	e := New()
	e.Use(Logger())
	e.Use(Recovery())
	return e
}

func Init() HandlerFunc {
	return func(c *Context) {
		// auth
		s := c.GetHeader(consts.Authorization)
		if s == "" {
			s, _ = c.GetQuery(consts.Authorization)
		}
		if s != "" {
			c.Set(consts.CTokenRaw, s)

			ss := strings.Split(s, " ")
			if len(ss) > 1 {
				s = ss[len(ss)-1]
			}

			t, err := token.Parse(s)
			if err == nil {
				c.Set(consts.CToken, t)
			}
		}

		// trace_id
		tid := c.GetHeader(consts.TraceId)
		if tid == "" {
			tid = seq.NewUUID().String()
			c.Header(consts.TraceId, tid)
		}
		c.Set(consts.CTraceId, tid)
	}
}

func PrintRoutes(e *Engine) {
	routes := e.Routes()
	sort.Slice(routes, func(i, j int) bool { return routes[i].Path < routes[j].Path })

	ml := 0
	for i := 0; i < len(routes); i++ {
		l := len(routes[i].Path)
		if l > ml {
			ml = l
		}
	}
	ml += 5

	sb := strings.Builder{}
	sb.WriteString("serving path list\n")
	for i := 0; i < len(routes); i++ {
		sb.WriteString(fmt.Sprintf("  %-7s %-"+strconv.Itoa(ml)+"s %s\n", routes[i].Method, routes[i].Path, routes[i].Handler))
	}

	log.Debug(sb.String())
}
