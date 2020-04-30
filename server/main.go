package main

import (
	"fmt"
	database "server/database"
	"server/router"
)

func main() {
	fmt.Println("start server")
	defer database.InitGORMDb().Close()

	router := router.InitRouter()
	router.Run(":8880")

}
