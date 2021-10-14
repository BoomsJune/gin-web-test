package main

import (
	"flag"

	"example.com/web-test/cmd/router"
	"example.com/web-test/config"
	"example.com/web-test/internal/pkg/db/migrate"
)

var migrateDB bool

func init() {
	flag.BoolVar(&migrateDB, "db", false, "migrate all db")
}

func main() {
	flag.Parse()
	if migrateDB {
		migrate.MigrateAll()
	}

	engine := router.Register()
	engine.Run(config.Cfg.App.Listen)
}
