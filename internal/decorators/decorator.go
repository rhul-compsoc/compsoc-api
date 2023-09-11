package decorators

import "github.com/gin-gonic/gin"

type DecoratorFunc func(fn gin.HandlerFunc) gin.HandlerFunc
