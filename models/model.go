package models

type Books struct {
	ID          int    `json:"id" form:"id" gorm:"primaryKey"`
	Title       string `json:"title" form:"title" binding:"required"`
	Author      string `json:"author" form:"author" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Stock       int    `json:"stock" form:"stock" binding:"required"`
}

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

//validate username
const (
	USER     = "admin"
	PASSWORD = "123"
	SECRET   = "secret"
)
