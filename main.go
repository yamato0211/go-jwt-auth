package main

import (
	"fmt"
	"jwt-tutorial/db"
	"jwt-tutorial/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	utils.LoadEnv()
	db.InitDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messaga": "Hello Golang!",
		})
	})

	r.Run(fmt.Sprintf(":%s", utils.ApiPort))
}
