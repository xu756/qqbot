package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlDB() {
	//dsn := "xu:xjx756756@tcp(121.5.132.57:5700)/ginvue2?charset=utf8&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	DB = db
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("用户表创建失败")
		return
	}
}
