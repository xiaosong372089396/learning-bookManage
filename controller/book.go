package controller

import (
	"bookManage/dao/mysql"
	"bookManage/model"

	"github.com/gin-gonic/gin"
)

// 添加数据
func CreateBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	mysql.DB.Create(p)
	c.JSON(200, gin.H{"msg": "success"})
}

// 查看书籍列表
func GetBookListHandler(c *gin.Context) {
	books := []model.Book{}
	mysql.DB.Find(&books)
	c.JSON(200, gin.H{"books": books})
}
