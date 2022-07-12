package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlDB() {
	dsn := "xu:xjx756756@tcp(0.0.0.0:4700)/bot?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	DB = db
	err = DB.AutoMigrate(&User{}, &UserIntegral{})
	if err != nil {
		fmt.Println("用户表创建失败")
		return
	}
}
