package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("/static", false)))
	// Handle 404s by serving the index.html page
	router.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})
	router.Run()
}
