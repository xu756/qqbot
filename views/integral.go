package views

import (
	"bot/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
)

func SignIn(data Data, ws *websocket.Conn) {
	var db = models.DB
	var user models.User
	db.Where("user_id = ?", data.UserId).First(&user)
	//用户不存在
	if user.Id == 0 {
		user.UserId = data.UserId
		user.Integral = 100
		user.Permission = 0
		db.Create(&user)
	}
	var num = rand.Int63n(100)
	user.Integral += num
	db.Save(&user)
	res := make(map[string]interface{})
	if num == 100 {
		res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n签到成功\n恭喜中大奖100金币\n目前金币%d枚", user.UserId, user.Integral)
	}
	res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n签到成功\n获得%d金币\n目前金币%d枚", user.UserId, num, user.Integral)
	res["group_id"] = data.GroupId
	ws.WriteJSON(gin.H{
		"action": "send_group_msg",
		"params": res,
	})

}
