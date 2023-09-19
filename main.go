package main

import (
	"github.com/Asad2730/anstAPI/connect"
	"github.com/gin-gonic/gin"
)

func init() {
	connect.Connect()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
