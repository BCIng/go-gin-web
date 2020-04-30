package model

import (
	"fmt"
	"log"
	"server/config"
	model "server/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitGORMDb() *gorm.DB {
	fmt.Println("connected database")
	c := config.GetConfig()
	db, err := gorm.Open("mysql", c.GetString("dbAdddress"))
	if err != nil {
		log.Fatal("Database connection failed. Database url: "+c.GetString("dbAdddress")+" error: ", err)
	} else {
		fmt.Print("\n\n------------------------------------------ GORM OPEN SUCCESS! -----------------------------------------------\n\n")
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;").AutoMigrate()
	db.AutoMigrate(&model.User{})

	db.LogMode(c.GetBool("dblog"))
	DB = db

	return db

}
