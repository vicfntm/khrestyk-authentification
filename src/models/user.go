package models

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required" `
}

type LoginUserStruct struct {
	Id       int    `json:"-"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required" `
}
