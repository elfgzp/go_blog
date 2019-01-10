package main

import (
	"fmt"
	"github.com/elfgzp/go_blog/controllers"
	"github.com/elfgzp/go_blog/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	db := models.ConnectToDB()
	defer db.Close()
	models.SetDB(db)

	controllers.Startup()
	port := 8808
	fmt.Printf("server run at http://127.0.0.1:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
