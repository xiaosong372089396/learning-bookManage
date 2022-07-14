package main

import (
	"bookManage/dao/mysql"
	"bookManage/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化mysql
	mysql.InitMysql()
	fmt.Println(mysql.DB, 1111111)

	// 1: 将实例化router服务的方法拆分到router文件夹下
	r := router.InitRouter()

	// 2: 定义路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "ok")
	})

	// 3: 启动服务
	r.Run(":8000")
}
