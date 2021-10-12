package main

import (
	"example.com/web-test/cmd/myapp/migrate"
	"example.com/web-test/cmd/myapp/router"
	"example.com/web-test/internal/pkg/config"
)

func main() {
	migrate.MigrateAll()

	engine := router.Register()
	engine.Run(config.Cfg.App.Listen)
}
