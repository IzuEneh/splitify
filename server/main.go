package main

import (
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/splitify/controller"
)

func main() {
	port := os.Getenv("PORT")

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./static", false)))
	router.POST("/generate", controller.GeneratePlaylists)
	router.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	router.Run(":" + port)
}
