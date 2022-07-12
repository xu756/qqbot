package views

import (
	"bot/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
)

func SignIn(data Data, ws *websocket.Conn) {
	var db = models.DB
	var user models.User
	var userIntegral models.UserIntegral //签到表
	userIntegral.UserId = data.UserId
	var now = time.Now().Format("2006-01-02")
	db.Where("user_id = ?", data.UserId).First(&user)
	db.Where("user_id = ? and date = ?", data.UserId, now).First(&userIntegral)
	//用户不存在
	if user.Id == 0 {
		user.UserId = data.UserId
		user.Integral = 100
		user.Permission = 0
		db.Create(&user)
		// 跳过签到
	}
	res := make(map[string]interface{})
	res["group_id"] = data.GroupId
	if userIntegral.Id == 0 {
		var num = rand.Int63n(100) //生成签到金币
		userIntegral.GetIntegral = num
		userIntegral.Date = now
		user.Integral += num
		if num == 100 {
			res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n签到成功\n恭喜中大奖100金币\n目前金币%d枚", user.UserId, user.Integral)
		} else {
			res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n签到成功\n获得%d金币\n目前金币%d枚", user.UserId, num, user.Integral)
		}
		ws.WriteJSON(gin.H{
			"action": "send_group_msg",
			"params": res,
		})
		defer db.Save(&user)
		defer db.Create(&userIntegral)
	} else {
		res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n不要重复签到", user.UserId)
		ws.WriteJSON(gin.H{
			"action": "send_group_msg",
			"params": res,
		})
	}

}
