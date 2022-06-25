package router

import (
	"alta-test/controller/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine, u handler.HandlerUser) {
	c.Use(gin.Logger())
	c.Use(gin.Recovery())
	c.Use(cors.Default())

	user := c.Group("/user")
	user.POST("", u.CreateUser())
	user.GET("", u.GetAllUsers())
	user.GET("/:id", u.GetUserID())
	user.PUT("/:id", u.UpdateUserID())
	user.DELETE("/:id", u.DeleteUserID())
}
