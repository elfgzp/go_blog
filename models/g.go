package models

import (
	"fmt"
	"github.com/elfgzp/go_blog/config"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connect to db ...")
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		panic(fmt.Errorf("Failed to connect database %s", err))
	}
	db.SingularTable(true)
	return db
}
