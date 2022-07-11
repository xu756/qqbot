package views

import (
	"bot/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Sender struct {
	UserId   int64  `json:"user_id"`  //发送事件的用户的 QQ 号
	NickName string `json:"nickname"` //发送事件的用户的昵称
	Sex      string `json:"sex"`      //性别
	Age      int32  `json:"age"`      //年龄
	Card     string `json:"card"`     //群名片		//群
	Area     string `json:"area"`     //地区
	Level    string `json:"level"`    //群等级
	Role     string `json:"role"`     //群角色
	Title    string `json:"title"`    //群头衔

}

func Message(data Data, ws *websocket.Conn) {
	switch data.MessageType {
	case "private":
		ws.WriteJSON(gin.H{
			"action": "send_msg",
			"params": gin.H{
				"message_type": "private",
				"user_id":      "756334744",
				"message": gin.H{
					"type": "shake",
				},
			},
		})

	case "group":
		if data.Message == "签到" {
			models.InitMysqlDB()
			var db = models.DB
			var user models.User
			user.UserId = data.UserId
			db.Where("user_id = ?", data.UserId).First(&user)
			println(user.UserId)
			if user.Id == 0 {
				user.UserId = data.UserId
				user.Integral = 100
				user.Permission = 0
				db.Create(&user)
				return
			}
			if user.Integral == 0 {
				ws.WriteJSON(gin.H{
					"action": "send_msg",
					"params": gin.H{
						"message_type": "group",
						"group_id":     data.GroupId,
						"message":      fmt.Sprintf("%s(%d) 金币不足 请别人送金币", data.Sender.Card, data.Sender.UserId),
					},
				})
				return
			}
			ws.WriteJSON(gin.H{
				"action": "send_msg",
				"params": gin.H{
					"message_type": "group",
					"group_id":     data.GroupId,
					"message":      "签到成功",
				},
			})
			return
		}
	}

}