package migrate

import (
	"log"

	"example.com/web-test/internal/myapp/auth"
	"example.com/web-test/internal/pkg/db"
)

func MigrateAll() {
	db.DB.AutoMigrate(
		&auth.User{},
	)

	log.Println("Migrate db.")
}
