package db

import (
	"log"
	"time"

	"example.com/web-test/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	maxIdleConns = 20        // 空闲最大连接数
	maxOpenConns = 200       // 最大打开连接数
	maxLifetime  = time.Hour // 连接可复用最大时间
	maxIdletime  = time.Hour // 空闲连接可复用最大时间
)

var DB *gorm.DB

func init() {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: config.Cfg.DB.Url,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名去掉s
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	} else {
		db, err := conn.DB()
		if err != nil || db.Ping() != nil {
			panic(err)
		} else {
			DB = conn
			db.SetMaxIdleConns(maxIdleConns)
			db.SetMaxOpenConns(maxOpenConns)
			db.SetConnMaxLifetime(maxLifetime)
			db.SetConnMaxIdleTime(maxIdletime)
			log.Println("Postgresql has already prepared for user connection.")
		}
	}
}
