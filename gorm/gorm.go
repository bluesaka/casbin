package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"))
	if err != nil {
		log.Fatal(err)
	}

	mysqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// 最大打开的连接数
	mysqlDB.SetMaxOpenConns(100)

	// 最大闲置连接数
	mysqlDB.SetMaxIdleConns(50)

	// 最大连接时间
	mysqlDB.SetConnMaxLifetime(3600 * time.Second)

	// 最大闲置时间
	mysqlDB.SetConnMaxIdleTime(3600 * time.Second)

	DB = db
}
