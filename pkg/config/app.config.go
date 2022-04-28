package config

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	sqlDB, err := sql.Open("mysql", "books_db")
	if err != nil {
		panic(err)
	}
	conn, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}
