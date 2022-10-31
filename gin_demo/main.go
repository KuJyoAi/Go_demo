package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BotMessage struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
	Type    int64  `json:"type"`
}

func main() {
	r := gin.Default()
	r.POST("/", func(ctx *gin.Context) {
		meg := &BotMessage{}
		ctx.ShouldBindJSON(meg)
		fmt.Println(meg)
	})
	r.GET("/", func(ctx *gin.Context) {
		fmt.Println("GET")
	})
	r.Run()
}
