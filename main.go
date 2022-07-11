package main

import (
	"bot/views"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//websocket 路由
	r.GET("/qq", views.Bot)
	//监听端口默认为8080
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
