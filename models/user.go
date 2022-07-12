package models

import "time"

type User struct {
	Id         int64 `primary_key:"true"`
	UserId     int64 //发送事件的用户的 QQ 号
	Integral   int64 //积分
	Permission int64 //权限
}
type UserIntegral struct {
	Id          int64 `primary_key:"true"`
	UserId      int64 //发送事件的用户的 QQ 号
	GetIntegral int64 //积分
	CreatedAt   time.Time
	Date        string //日期
}
