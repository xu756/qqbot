package main

import (
	"bot/views"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//websocket 路由
	r.GET("/qq", views.Bot)
	err := r.Run(":7567")
	if err != nil {
		return
	}
}
