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
	log.Println("DB Init ...")
	db := models.ConnectToDB()
	defer db.Close()
	models.SetDB(db)

	db.DropTableIfExists(models.User{}, models.Post{}, "follower")
	db.CreateTable(models.User{}, models.Post{})

	models.AddUser("elfgzp", "abc123", "me@elfgzp.im")
	models.AddUser("jerry", "abc123", "rene@test.com")

	u1, _ := models.GetUserByUsername("elfgzp")
	u1.CreatePost("Beautiful day in Portland!")
	models.UpdateAboutMe(u1.Username, `I am an engineer.`)

	u2, _ := models.GetUserByUsername("jerry")
	u2.CreatePost("The Avengers movie was so cool!")
	u2.CreatePost("Sun shine is beautiful")
	models.UpdateAboutMe(u1.Username, `I am Jerry.`)

	u1.Follow(u2.Username)
}
