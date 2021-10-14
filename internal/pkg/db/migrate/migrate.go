package migrate

import (
	"log"

	"example.com/web-test/internal/auth"
	"example.com/web-test/internal/pkg/db"
)

var models = []interface{}{
	&auth.User{},
}

func MigrateAll() {
	if err := db.DB.AutoMigrate(models...); err != nil {
		panic(err)
	}
	log.Println("Migrate db.")
}

func DropAll() {
	if err := db.DB.Migrator().DropTable(models...); err != nil {
		panic(err)
	}

	log.Println("Drop all tables.")
}
