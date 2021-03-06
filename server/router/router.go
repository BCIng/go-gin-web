package router

import (
	. "server/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/register", Register)
	router.POST("/register", Register)

	return router

}
