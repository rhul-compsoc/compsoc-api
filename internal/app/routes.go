package app

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/decorators"
	"github.com/rhul-compsoc/compsoc-api-go/internal/handlers"
	"github.com/rhul-compsoc/compsoc-api-go/internal/router"
)

// All API routes
var routes = router.Routes{
	RouteInfo: []router.Route{
		// Ping Handlers
		{
			Name:        "Get Ping",
			Method:      router.Get,
			Path:        "/ping",
			Params:      "",
			HandlerFunc: handlers.PingGet(s),
		},
		// Admin Test Handler
		{
			Name:        "Get Admin Ping",
			Method:      router.Get,
			Path:        "/ping/admin",
			Params:      "",
			HandlerFunc: handlers.AdminPingGet(),
			DecoratorFunc: decorators.AdminOnly(),
		},
		// Guild Handlers
		{
			Name:        "Get Guild",
			Method:      router.Get,
			Path:        "/guild",
			Params:      "/:guild",
			HandlerFunc: handlers.GuildGet(),
		},
		// Member Handlers
		{
			Name:        "List Member",
			Method:      router.Get,
			Path:        "/member",
			Params:      "",
			HandlerFunc: handlers.MemberList(s),
		},
		{
			Name:        "Get Member",
			Method:      router.Get,
			Path:        "/member",
			Params:      "/:member",
			HandlerFunc: handlers.MemberGet(s),
		},
		{
			Name:        "Post Member",
			Method:      router.Post,
			Path:        "/member",
			Params:      "",
			HandlerFunc: handlers.MemberPost(s),
		},
		{
			Name:        "Put Member",
			Method:      router.Put,
			Path:        "/member",
			Params:      "",
			HandlerFunc: handlers.MemberPut(s),
		},
		{
			Name:        "Patch Member",
			Method:      router.Patch,
			Path:        "/member",
			Params:      "",
			HandlerFunc: handlers.MemberPatch(s),
		},
		{
			Name:        "Delete Member",
			Method:      router.Delete,
			Path:        "/member",
			Params:      "/:member",
			HandlerFunc: handlers.MemberDelete(s),
		},
		// User Handlers
		{
			Name:        "List User",
			Method:      router.Get,
			Path:        "/user",
			Params:      "",
			HandlerFunc: handlers.UserList(s),
		},
		{
			Name:        "Get User",
			Method:      router.Get,
			Path:        "/user",
			Params:      "/:user",
			HandlerFunc: handlers.UserGet(s),
		},
		{
			Name:        "Post User",
			Method:      router.Post,
			Path:        "/user",
			Params:      "",
			HandlerFunc: handlers.UserPost(s),
		},
		{
			Name:        "Put User",
			Method:      router.Put,
			Path:        "/user",
			Params:      "",
			HandlerFunc: handlers.UserPut(s),
		},
		{
			Name:        "Patch User",
			Method:      router.Patch,
			Path:        "/user",
			Params:      "",
			HandlerFunc: handlers.UserPatch(s),
		},
		{
			Name:        "Delete User",
			Method:      router.Delete,
			Path:        "/user",
			Params:      "/:user",
			HandlerFunc: handlers.UserDelete(s),
		},
		// Event Handlers
		{
			Name:        "List Event",
			Method:      router.Get,
			Path:        "/event",
			Params:      "",
			HandlerFunc: handlers.EventList(s),
		},
		{
			Name:        "Get Event",
			Method:      router.Get,
			Path:        "/event",
			Params:      "/:event",
			HandlerFunc: handlers.EventGet(s),
		},
		{
			Name:        "Post Event",
			Method:      router.Post,
			Path:        "/event",
			Params:      "",
			HandlerFunc: handlers.EventPost(s),
		},
		{
			Name:        "Put Event",
			Method:      router.Put,
			Path:        "/event",
			Params:      "",
			HandlerFunc: handlers.EventPut(s),
		},
		{
			Name:        "Patch Event",
			Method:      router.Patch,
			Path:        "/event",
			Params:      "",
			HandlerFunc: handlers.EventPatch(s),
		},
		{
			Name:        "Delete Event",
			Method:      router.Delete,
			Path:        "/event",
			Params:      "/:event",
			HandlerFunc: handlers.EventDelete(s),
		},
		// File Controllers
		{
			Name:        "List File",
			Method:      router.Get,
			Path:        "/file",
			Params:      "",
			HandlerFunc: handlers.ListFile(),
		},
		{
			Name:        "Get File",
			Method:      router.Get,
			Path:        "/file",
			Params:      "/:file",
			HandlerFunc: handlers.GetFile(),
		},
		{
			Name:        "Post File",
			Method:      router.Post,
			Path:        "/file",
			Params:      "",
			HandlerFunc: handlers.PostFile(),
		},
		{
			Name:        "Post Files",
			Method:      router.Post,
			Path:        "/files",
			Params:      "",
			HandlerFunc: handlers.PostFiles(),
		},
	},
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
