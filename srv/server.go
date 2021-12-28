package srv

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/go-sdk/lib/conf"
	"github.com/go-sdk/lib/log"
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
	return gin.New()
}

func Default() *Engine {
	e := New()
	e.Use(Logger())
	e.Use(Recovery())
	return e
}

func PrintRoutes(e *Engine) {
	routes := e.Routes()
	sort.Slice(routes, func(i, j int) bool { return routes[i].Path < routes[j].Path })
	for i := 0; i < len(routes); i++ {
		route := routes[i]
		log.Debugf("%-7s %-37s %s", route.Method, route.Path, route.Handler)
	}
}
