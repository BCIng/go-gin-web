package main

import (
	"fmt"
	"server/models"
	"server/routers"
)

func main() {
	fmt.Println("start server")
	defer models.InitDb().Close()
	router := routers.InitRouter()
	router.Run(":8880")
}
