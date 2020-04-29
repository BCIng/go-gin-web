package models

import (
	"github.com/jinzhu/gorm"
)

type Register struct {
	gorm.Model
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}

// insert
func (register Register) Insert() (err error) {
	result := db.Create(&register)

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//getUserName
// func (unserName string) GetUserName(register Register, err error) {
// 	db.Where("username = ?", useName).First(&register).Error
// }
