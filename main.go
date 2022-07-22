package main

import (
	"jwt-tutorial/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messaga": "Hello Golang!",
		})
	})

	// func Private(){
	// 	post := &post{
	// 		Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
	// 		Tag:   "Go",
	// 		URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
	// 	}
	// 	json.NewEncode(w).Encode(post)
	// }

	r.GET("/auth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"token": auth.GenerateToken,
		})
	})

	//r.GET("/private", auth.JwtMiddleware.)

	r.Run(":8000")
}
