package views

import (
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
				"user_id": "756334744",
				"message": gin.H{
					"type": "face",
					"data": gin.H{
						"id": "13",
					},
				},
			},
		})

	case "group":
		switch data.Message {
		case "签到":
			SignIn(data, ws)
		case "抽金币":
			DrawGold(data, ws)
		case "个人信息":
			GetUserInfo(data, ws)
		}

	}

}
