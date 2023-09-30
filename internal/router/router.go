package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/pkg/util"
)

// Router stores pointers of both a gin engine and a store.
//   - Engine : gin engine that is responsible for routing.
type Router struct {
	Engine *gin.Engine
}

// Returns a pointer to a default router.
//   - Engine : default gin router.
func New() Router {
	return Router{
		Engine: gin.Default(),
	}
}

// Wrapper for the gin engine NoRoute method.
func (r *Router) NoRoute(f gin.HandlerFunc) {
	r.Engine.NoRoute(f)
}

// Wrapper for the gin engine Use method.
func (r *Router) Use(f gin.HandlerFunc) {
	r.Engine.Use(f)
}

// Wrapper for the gin engine Run method.
func (r *Router) Run() {
	r.Engine.Run()
}

func (r *Router) Close() {

}

// Wrapper for the gin engine GET method.
func (r *Router) Get(route Route) {
	if route.DecoratorFunc == nil {
		r.Engine.GET(route.Path+route.Params, route.HandlerFunc)
	} else {
		r.Engine.GET(route.Path+route.Params, route.DecoratorFunc(route.HandlerFunc))
	}
}

// Wrapper for the gin engine POST method.
func (r *Router) Post(route Route) {
	if route.DecoratorFunc == nil {
		r.Engine.POST(route.Path+route.Params, route.HandlerFunc)
	} else {
		r.Engine.POST(route.Path+route.Params, route.DecoratorFunc(route.HandlerFunc))
	}
}

// Wrapper for the gin engine PUT method.
func (r *Router) Put(route Route) {
	if route.DecoratorFunc == nil {
		r.Engine.PUT(route.Path+route.Params, route.HandlerFunc)
	} else {
		r.Engine.PUT(route.Path+route.Params, route.DecoratorFunc(route.HandlerFunc))
	}
}

// Wrapper for the gin engine PATCH method.
func (r *Router) Patch(route Route) {
	if route.DecoratorFunc == nil {
		r.Engine.PATCH(route.Path+route.Params, route.HandlerFunc)
	} else {
		r.Engine.PATCH(route.Path+route.Params, route.DecoratorFunc(route.HandlerFunc))
	}
}

// Wrapper for the gin engine DELETE method.
func (r *Router) Delete(route Route) {
	if route.DecoratorFunc == nil {
		r.Engine.DELETE(route.Path+route.Params, route.HandlerFunc)
	} else {
		r.Engine.DELETE(route.Path+route.Params, route.DecoratorFunc(route.HandlerFunc))
	}
}

func (r *Router) Options(route Route) {
	if route.DecoratorFunc == nil {
		r.Engine.OPTIONS(route.Path+route.Params, route.HandlerFunc)
	} else {
		r.Engine.OPTIONS(route.Path+route.Params, route.DecoratorFunc(route.HandlerFunc))
	}
}

// Registers a route.
//
// If the route method is undefined it will cause a panic.
func (r *Router) RegisterRoute(route Route) {
	switch route.Method {
	case Undefined:
		util.ErrOut(util.ErrUndefinedRouteMethod)
	case Get:
		r.Get(route)
	case Post:
		r.Post(route)
	case Put:
		r.Put(route)
	case Patch:
		r.Patch(route)
	case Delete:
		r.Delete(route)
	case Options:
		r.Options(route)
	}
}

// Registers all routes passed depending on there method type.
//
// If the route method is undefined it will cause a panic.
func (r *Router) RegisterRoutes(routes Routes) {
	for _, route := range routes.RouteInfo {
		r.RegisterRoute(route)
	}
}
