package model

type User struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}
