package main

import (
	"example.com/web-test/cmd/myapp/router"
	"example.com/web-test/internal/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Register(r)
	r.Run(config.Cfg.App.Listen)
}
