package model

/*
  作用: 用户鱼书籍管理的关系表
*/

type BookUser struct {
	UserID int64 `gorm: "primaryKey"`
	BookID int64 `gorm: "primaryKey"`
}
