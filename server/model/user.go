package model

type User struct {
	BaseModel
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Ip       string `form:"ip" json:"ip" binding:"required"`
}
