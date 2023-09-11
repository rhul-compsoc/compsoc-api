package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/decorators"
)

// Route struct stores information for create a route in the gin engine.
//   - Name : stores the name of the route.
//   - Method : what type is it? Post, Patch, Delete, etc.?
//   - Path : path of the route.
//   - Params : parameters of the route.
//   - HandlerFunc : gin.Handlerfunc,
//     stores the method that occurs when this route is queried.
//   - DecoratorFunc: decorators.DecoratorFunc,
//     allows for the addition of auth and other types of decorator
type Route struct {
	Name          string                   // Route name
	Method        Method                   // Route method
	Path          string                   // Route path
	Params        string                   // Route Parameters
	HandlerFunc   gin.HandlerFunc          // Route Handler Function
	DecoratorFunc decorators.DecoratorFunc // Route Decorator Function
}

// Routes stores a slice of the Route struct.
//   - RouteInfo : slice of Route structs.
type Routes struct {
	RouteInfo []Route // Routes-
}
