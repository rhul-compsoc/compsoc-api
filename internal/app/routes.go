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
				Name:        "Get Ping",
				Method:      router.Get,
				Path:        "/ping",
				Handler:     "",
				HandlerFunc: handlers.PingGet(s),
			},
			// Guild Handlers
			{
				Name:        "Get Guild",
				Method:      router.Get,
				Path:        "/guild",
				Handler:     "/:guild",
				HandlerFunc: handlers.GuildGet(),
			},
			// Member Handlers
			{
				Name:        "List Member",
				Method:      router.Get,
				Path:        "/member",
				Handler:     "",
				HandlerFunc: handlers.MemberList(s),
			},
			{
				Name:        "Get Member",
				Method:      router.Get,
				Path:        "/member",
				Handler:     "/:member",
				HandlerFunc: handlers.MemberGet(s),
			},
			{
				Name:        "Post Member",
				Method:      router.Post,
				Path:        "/member",
				Handler:     "",
				HandlerFunc: handlers.MemberPost(s),
			},
			{
				Name:        "Put Member",
				Method:      router.Put,
				Path:        "/member",
				Handler:     "",
				HandlerFunc: handlers.MemberPut(s),
			},
			{
				Name:        "Patch Member",
				Method:      router.Patch,
				Path:        "/member",
				Handler:     "",
				HandlerFunc: handlers.MemberPatch(s),
			},
			{
				Name:        "Delete Member",
				Method:      router.Delete,
				Path:        "/member",
				Handler:     "/:member",
				HandlerFunc: handlers.MemberDelete(s),
			},
			// User Handlers
			{
				Name:        "List User",
				Method:      router.Get,
				Path:        "/user",
				Handler:     "",
				HandlerFunc: handlers.UserList(s),
			},
			{
				Name:        "Get User",
				Method:      router.Get,
				Path:        "/user",
				Handler:     "/:user",
				HandlerFunc: handlers.UserGet(s),
			},
			{
				Name:        "Post User",
				Method:      router.Post,
				Path:        "/user",
				Handler:     "",
				HandlerFunc: handlers.UserPost(s),
			},
			{
				Name:        "Put User",
				Method:      router.Put,
				Path:        "/user",
				Handler:     "",
				HandlerFunc: handlers.UserPut(s),
			},
			{
				Name:        "Patch User",
				Method:      router.Patch,
				Path:        "/user",
				Handler:     "",
				HandlerFunc: handlers.UserPatch(s),
			},
			{
				Name:        "Delete User",
				Method:      router.Delete,
				Path:        "/user",
				Handler:     "/:user",
				HandlerFunc: handlers.UserDelete(s),
			},
			// Event Handlers
			{
				Name:        "List Event",
				Method:      router.Get,
				Path:        "/event",
				Handler:     "",
				HandlerFunc: handlers.EventList(s),
			},
			{
				Name:        "Get Event",
				Method:      router.Get,
				Path:        "/event",
				Handler:     "/:event",
				HandlerFunc: handlers.EventGet(s),
			},
			{
				Name:        "Post Event",
				Method:      router.Post,
				Path:        "/event",
				Handler:     "",
				HandlerFunc: handlers.EventPost(s),
			},
			{
				Name:        "Put Event",
				Method:      router.Put,
				Path:        "/event",
				Handler:     "",
				HandlerFunc: handlers.EventPut(s),
			},
			{
				Name:        "Patch Event",
				Method:      router.Patch,
				Path:        "/event",
				Handler:     "",
				HandlerFunc: handlers.EventPatch(s),
			},
			{
				Name:        "Delete Event",
				Method:      router.Delete,
				Path:        "/event",
				Handler:     "/:event",
				HandlerFunc: handlers.EventDelete(s),
			},
			// File Controllers
			{
				Name:        "Get File",
				Method:      router.Get,
				Path:        "/file",
				Handler:     "/:file",
				HandlerFunc: handlers.GetFile(),
			},
			{
				Name:        "Post File",
				Method:      router.Post,
				Path:        "/file",
				Handler:     "",
				HandlerFunc: handlers.PostFile(),
			},
			{
				Name:        "Post Files",
				Method:      router.Post,
				Path:        "/files",
				Handler:     "",
				HandlerFunc: handlers.PostFiles(),
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
