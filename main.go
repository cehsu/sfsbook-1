package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", IndexHandler)
	router.GET("/search", SearchHandler)
	router.GET("/resources/:name", ResourceShowHandler)

	router.Run() // listen and serve on 0.0.0.0:8080
}
