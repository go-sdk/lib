package srv

import (
	"time"

	"github.com/gin-contrib/cors"
)

type CORSConfig = cors.Config

func CORS() HandlerFunc {
	return CORSWithConfig(CORSConfig{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
}

func CORSWithConfig(config CORSConfig) HandlerFunc {
	return cors.New(config)
}
