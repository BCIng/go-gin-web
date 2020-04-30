package repository

import (
	database "server/database"
	model "server/model"
)

func UserCreate(user model.User) error {
	return database.DB.Create(&user).Error
}

func UserRead(user model.User) error {
	return database.DB.Where("username =?", user.Username).First(&user).Error
}
