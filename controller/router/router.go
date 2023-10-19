package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RouteHandlers interface of handlers
type RouteHandlers interface {
	URLMapping(r *gin.RouterGroup)
}

// handlers register an endpoint with handler here.
// it will automatic registered into routers
var handlers = map[string]RouteHandlers{}

func Router() *gin.Engine {

	c := gin.New()
	c.Use(gin.Logger())
	c.Use(gin.Recovery())
	c.Use(cors.Default())

	v := c.Group("v1/")

	for p, h := range handlers {
		h.URLMapping(v.Group(p))
	}

	return c
}
