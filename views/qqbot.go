package views

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Data struct {
	Time        int64  `json:"time"`         //时间戳
	SelfId      int64  `json:"self_id"`      //收到事件的机器人的 QQ 号
	PostType    string `json:"post_type"`    //表示该上报的类型, 消息(message), 请求(request), 通知(notice), 或元事件(元事件)
	MessageType string `json:"message_type"` //表示消息的类型, 私聊(private), 群聊(group)
	UserId      int64  `json:"user_id"`      //发送事件的用户的 QQ 号
	GroupId     int64  `json:"group_id"`     //发送事件的群号
	MessageId   int32  `json:"message_id"`   //消息的 id
	Message     string `json:"message"`      //消息的内容
	Sender      Sender `json:"sender"`       //发送事件的用户的信息
}

func Bot(c *gin.Context) {
	ws, _ := upGrader.Upgrade(c.Writer, c.Request, nil)
	for {
		//读取json
		var data Data
		err := ws.ReadJSON(&data)
		if err != nil {
			log.Panicln("data 读取json错误")
		}

		switch data.PostType {
		case "message":
			Message(data, ws)
			//消息

		case "request":
			//请求
		case "notice":
			//通知
		case "meta_event":
			//元事件
		}

	}

}
