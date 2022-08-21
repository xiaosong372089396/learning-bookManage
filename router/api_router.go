package router

import (
	"bookManage/controller"
	"bookManage/middleware"

	"github.com/gin-gonic/gin"
)

func LoadAPIRouter(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)

	v1 := r.Group("/api/v1")
	// 添加中间验证
	v1.Use(middleware.AuthMiddleware())
	v1.POST("book", controller.CreateBookHandler)
	v1.GET("book", controller.GetBookListHandler)
	v1.GET("book/:id", controller.GetBookDetailHandler)
	v1.PUT("book", controller.UpdateBookHandler)
	v1.DELETE("book/:id", controller.DeleteBookHandler)
}
