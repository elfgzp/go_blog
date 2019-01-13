package main

import (
	"fmt"
	"github.com/elfgzp/go_blog/config"
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

	users := []models.User{
		{
			Username:     "elfgzp",
			PasswordHash: models.GeneratePasswordHash("abc123"),
			Email:        "me@elfgzp.com",
			Avatar:       fmt.Sprintf("%s/%s?d=identicon", config.GravatarURL, models.Md5("me@elfgzp.com")),
			Posts: []models.Post{
				{Body: "Beautiful day in Portland"},
			},
		},
		{
			Username:     "jerry",
			PasswordHash: models.GeneratePasswordHash("abc123"),
			Email:        "jerry@test.com",
			Avatar:       fmt.Sprintf("%s/%s?d=identicon", config.GravatarURL, models.Md5("jerry@test.com")),
			Posts: []models.Post{
				{Body: "The Avengers movie was so cool!"},
				{Body: "Sun shine is beautiful"},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}
}
