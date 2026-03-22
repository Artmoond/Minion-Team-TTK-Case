package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/local/file", func(c *gin.Context) {
		c.File("./assets/NITSHEMERTW.mp3")
	})

	router.Run(":8080")
}
