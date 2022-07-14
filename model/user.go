package model

/*
·json: "username" 定义json反向解析的名字·
gorm: "not null" 定义字段在数据库中不能为空
binding: "required" 请求时不能为空(shouldBind 参数校验)
*/
type User struct {
	Id       int64  `json:"id" gorm: "primaryKey"`
	Username string `json:"username" gorm: "not null" binding: "required"`
	Password string `json:"password" gorm: "not null" binding: "required"`
	Token    string `json:"token"`
}

// 表名默认添加s， 自定义表的名字
func (User) TableName() string {
	return "user"
}
