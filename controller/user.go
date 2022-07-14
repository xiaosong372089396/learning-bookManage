package controller

import (
	"bookManage/dao/mysql"
	"bookManage/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 注册功能
func RegisterHandler(c *gin.Context) {
	p := new(model.User)
	// 参数校验，绑定
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	// 入库: 最简单的创建数据
	mysql.DB.Create(p)
	c.JSON(200, gin.H{"msg": "success"})
}

// 登陆功能
func LoginHandler(c *gin.Context) {
	p := new(model.User)
	// 参数校验
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// 判断用户的用户名和密码是否正确
	u := model.User{Username: p.Username, Password: p.Password}
	// rows 判断是否有数据，nil证明没有这个用户
	if rows := mysql.DB.Where(&u).First(&u).Row(); rows == nil {
		c.JSON(403, gin.H{"msg": "密码错误"})
		return
	}
	// 随机生成一个字符串作为token
	token := uuid.New().String()
	mysql.DB.Model(u).Update("token", token)
	c.JSON(200, gin.H{"msg": "success", "token": token})
}
