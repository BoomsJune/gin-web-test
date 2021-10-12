package router

import (
	"net/http"

	"example.com/web-test/cmd/myapp/router/handler"
	"example.com/web-test/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.POST("/regist", handler.Register)
	r.POST("/login", handler.Login)

	authorized := r.Group("")
	authorized.Use(middleware.JwtAuthMiddleware())
	{
		authorized.GET("/info/:id", handler.Info)
	}
	return r
}
