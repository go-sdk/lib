package srv

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/go-sdk/lib/conf"
)

type (
	Engine  = gin.Engine
	Context = gin.Context

	HandlerFunc = gin.HandlerFunc
)

func init() {
	gin.SetMode(strings.ToLower(conf.Get("srv.gin.mode").StringD(gin.ReleaseMode)))
}

func New() *Engine {
	return gin.New()
}

func Default() *Engine {
	e := gin.New()
	e.Use(Logger())
	e.Use(Recovery())
	return e
}
