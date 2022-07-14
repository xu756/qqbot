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
		rand.Seed(time.Now().UnixNano())
		var num = rand.Int63n(50) //生成签到金币
		userIntegral.GetIntegral = num
		userIntegral.Date = now
		user.Integral += num
		if num == 100 {
			res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n签到成功\n恭喜中大奖100金币\n目前金币%d枚", user.UserId, user.Integral)
			user.Lucky += 1
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

// DrawGold 抽金币
func DrawGold(data Data, ws *websocket.Conn) {
	var db = models.DB
	var user models.User
	db.Where("user_id = ?", data.UserId).First(&user)
	res := make(map[string]interface{})
	res["group_id"] = data.GroupId
	if user.Id == 0 {
		user.UserId = data.UserId
		user.Integral = 100
		user.Permission = 0
		db.Create(&user)
	}
	if user.Integral < 50 {
		res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n您的金币(%d)不足10枚，无法抽奖", user.UserId, user.Integral)
		ws.WriteJSON(gin.H{
			"action": "send_group_msg",
			"params": res,
		})
		return
	}
	rand.Seed(time.Now().UnixNano())
	var num1 = rand.Int63n(20) + 2 //生成签到金币 要花费的
	var num2 = rand.Int63n(40)     //生成签到金币 获得的
	user.Integral = user.Integral - num1 + num2
	defer db.Save(&user)
	//中大奖
	if num2 == 40 {
		res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n恭喜中大奖40金币\n目前金币%d枚", user.UserId, user.Integral)
		ws.WriteJSON(gin.H{
			"action": "send_group_msg",
			"params": res,
		})
	}
	res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n抽奖成功\n消费%d枚金币\n获得%d枚金币\n目前金币%d枚", user.UserId, num1, num2, user.Integral)
	ws.WriteJSON(gin.H{
		"action": "send_group_msg",
		"params": res,
	})
}
func GetUserInfo(data Data, ws *websocket.Conn) {
	var db = models.DB
	var user models.User
	db.Where("user_id = ?", data.UserId).First(&user)
	res := make(map[string]interface{})
	res["group_id"] = data.GroupId
	if user.Id == 0 {
		user.UserId = data.UserId
		user.Integral = 100
		user.Permission = 0
		db.Create(&user)

	}
	res["message"] = fmt.Sprintf("[CQ:at,qq=%d]\n金币: %d 枚\n幸运值：%d\n罪恶值: %d", user.UserId, user.Integral, user.Lucky, user.Sin)
	ws.WriteJSON(gin.H{
		"action": "send_group_msg",
		"params": res,
	})

}
