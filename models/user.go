package models

type User struct {
	Id         int64 //用户id
	UserId     int64 //发送事件的用户的 QQ 号
	Integral   int64 //积分
	Permission int64 //权限
}
