package middleware

import (
	"bookManage/dao/mysql"
	"bookManage/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		var u model.User
		// 如果没有当前用户
		row := mysql.DB.Where("token = ?", token).First(&u).RowsAffected
		if row != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "当前token错误",
			})
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("UserId", u.Id)
		c.Next()
	}
}
