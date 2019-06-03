package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 初期値
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(context *gin.Context) {
		println("hello,Voters!!")
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}