package repository

import (
	database "server/database"
	model "server/model"
)

func UserCreate(user model.User) (err error) {
	result := database.DB.Create(&user)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
