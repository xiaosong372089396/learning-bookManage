package mysql

import (
	"bookManage/model"
	"fmt"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 为什么要讲变量定义全局， 是因为在其他地方都需要使用
var DB *gorm.DB

func InitMysql() {
	// 1、连接数据库
	dsn := "root:1@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("初始化mysql连接错误", err)
	}
	DB = db

	// 自动创建表结构
	if err := DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Println("自动创建表结构失败:", err)
	}
}
