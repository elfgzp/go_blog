package main

import (
	"github.com/elfgzp/go_blog/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


/*
create database go_blog_db  character set 'utf8mb4' COLLATE utf8mb4_unicode_ci;
 */
func main() {
	db := models.ConnectToDB()
	log.Println("DB Init ...")
	defer db.Close()
	models.SetDB(db)

	db.DropTableIfExists(models.User{}, models.Post{})
	db.CreateTable(models.User{}, models.Post{})
}
