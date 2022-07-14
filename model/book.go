package model

type Book struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding: "required"`
	Desc string `json:"desc" binding: "required"`
	User []User `gorm: "many2many:book_users"`
}

func (Book) TableName() string {
	return "book"
}
