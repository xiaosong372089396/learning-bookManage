package controller

import (
	"bookManage/dao/mysql"
	"bookManage/model"
	"strconv"

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

/*
// plan 2
type BookParams struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	UserId int64  `json:"user_id"`
}

// plan 2
func CreateBookUserPlan2Handler(c *gin.Context) {
	p := new(BookParams)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	// 用户添加关联关系
	book := model.Book{
		Id:   0,
		Name: p.Name,
		Desc: p.Desc,
		User: []model.User{
			{Id: p.UserId},
		},
	}
	mysql.DB.Create(&book)
	c.JSON(200, gin.H{"msg": "success"})
}
*/

// 添加数据和用户
func CreateBookUserHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	// 用户添加关联关系
	book := model.Book{
		Id:   0,
		Name: p.Name,
		Desc: p.Desc,
		User: []model.User{
			{Id: 4},
		},
	}
	mysql.DB.Create(&book)
	c.JSON(200, gin.H{"msg": "success"})
}

// 查看书籍列表
func GetBookListHandler(c *gin.Context) {
	books := []model.Book{}
	mysql.DB.Find(&books)
	c.JSON(200, gin.H{"books": books})
}

// 查看置顶书籍 http://127.0.0.1:8000/book/4/
func GetBookDetailHandler(c *gin.Context) {
	bookIdStr := c.Param("id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	book := model.Book{Id: bookId}
	mysql.DB.Find(&book)
	c.JSON(200, gin.H{"book": book})
}

// 修改操作
func UpdateBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	oldBook := &model.Book{Id: p.Id}
	var newBook model.Book
	if p.Name != "" {
		newBook.Name = p.Name
	}
	if p.Desc != "" {
		newBook.Desc = p.Desc
	}
	mysql.DB.Model(&oldBook).Updates(newBook)
	c.JSON(200, gin.H{"book": newBook})
}

// 删除
func DeleteBookHandler(c *gin.Context) {
	// 获取url参数
	pipelineIdStr := c.Param("id")
	bookId, _ := strconv.ParseInt(pipelineIdStr, 10, 64)
	// 删除book时， 也删除第三张表中的 用户对应关系记录
	mysql.DB.Select("Users").Delete(&model.Book{Id: bookId})
	c.JSON(200, gin.H{"msg": "success"})
}
