package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "books",
	}

	database, err := gorm.Open("mysql", cfg.FormatDSN())
	if err != nil {
    log.Fatal(err)
		panic(err)
	}
	db = database

	fmt.Println("Connected", cfg.FormatDSN())
}

func GetDB() *gorm.DB {
	return db
}
