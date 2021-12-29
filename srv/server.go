package srv

import (
	"fmt"
	"sort"
	"strconv"
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
