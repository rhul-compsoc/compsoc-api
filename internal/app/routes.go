package app

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/handlers"
	"github.com/rhul-compsoc/compsoc-api-go/internal/router"
)

// The makeRoutes method returns a pointer to a Routes struct,
// which stores a slice of Route.
func makeRoutes() router.Routes {
	return router.Routes{
		RouteInfo: []router.Route{
			// Ping Handlers
			{
				Name:        "GetPing",
				Method:      router.Get,
				Path:        "/ping",
				Handler:     "",
				HandlerFunc: handlers.PingGet(),
			},
			// Guild Handlers
			{
				Name:        "GetGuild",
				Method:      router.Get,
				Path:        "/guild",
				Handler:     "/:guild",
				HandlerFunc: handlers.GuildGet(),
			},
		},
	}
}

func reverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, _ := url.Parse("http://localhost:3000")
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL = c.Request.URL
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
