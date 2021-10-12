package db

import (
	"database/sql"
	"log"
	"time"

	"example.com/web-test/internal/pkg/config"
)

const (
	maxIdleConns = 20        // 空闲最大连接数
	maxOpenConns = 200       // 最大打开连接数
	maxLifetime  = time.Hour // 连接可复用最大时间
	maxIdletime  = time.Hour // 空闲连接可复用最大时间
)

var DB *sql.DB

func init() {
	db, err := sql.Open("postgres", config.Cfg.DB.Url)
	if err != nil {
		panic(err)
	} else {
		if err = db.Ping(); err != nil {
			panic(err)
		} else {
			db.SetMaxIdleConns(maxIdleConns)
			db.SetMaxOpenConns(maxOpenConns)
			db.SetConnMaxLifetime(maxLifetime)
			db.SetConnMaxIdleTime(maxIdletime)
			log.Println("Postgresql has already prepared for user connection.")
		}
	}
}
