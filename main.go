package main

import (
	"fmt"
	"github.com/elfgzp/go_blog/controllers"
	"net/http"
)

func main() {
	controllers.Startup()
	port := 8808
	fmt.Printf("server run at http://127.0.0.1:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%6d", port), nil)
}
