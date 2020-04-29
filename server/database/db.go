package database

import (
	"fmt"
	"log"
	"server/config"
	"server/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	fmt.Println("connected database")

	c := config.GetConfig()

	var err error

	db, err = gorm.Open("mysql", c.GetString("dbAdddress"))
	checkErr(err, "sql.Open failed")
	db.LogMode(c.GetBool("dblog"))

	db.AutoMigrate(&models.Register{})

}
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func InitDb() *gorm.DB {
	return db
}
