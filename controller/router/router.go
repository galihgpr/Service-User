package router

import (
	"alta-test/controller/handler"
	"alta-test/controller/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine, u handler.HandlerUser) {
	c.Use(gin.Logger())
	c.Use(gin.Recovery())
	c.Use(cors.Default())
	// Login
	c.POST("/login", u.Login())
	// Init Group
	user := c.Group("/user")
	// Use Middleware JWT
	user.Use(middlewares.MiddlewareJWT())
	{
		user.POST("", u.CreateUser())
		user.GET("", u.GetAllUsers())
		user.GET("/:id", u.GetUserID())
		user.PUT("/:id", u.UpdateUserID())
		user.DELETE("/:id", u.DeleteUserID())
	}
}
