package router

import (
	"net/http"

	"example.com/web-test/cmd/myapp/router/handler"
	"example.com/web-test/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	// engine.Use(gin.BasicAuth(gin.Accounts{
	// 	"admin": "admin",
	// }))
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	engine.POST("/regist", handler.Register)
	engine.POST("/login", handler.Login)

	authorized := engine.Group("")
	authorized.Use(middleware.JwtAuthMiddleware())
	{
		authorized.GET("/info/:id", handler.Info)
	}

}
