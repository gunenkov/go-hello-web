package main

import (
	firstpak "firstmodule/firstpak"
	secondpack "firstmodule/secondpak"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi")
	firstpak.Test()
	secondpack.SecondTest()

	r := gin.Default()
  	r.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
		"title": "hello-app",
		"version": 1,
		})
  	})
  	r.Run()
}