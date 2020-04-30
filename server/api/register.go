package api

import (
	"fmt"
	"net/http"
	"server/model"
	"server/repository"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerValue model.User
	registerValue.Username = c.Request.FormValue("username")
	registerValue.Password = c.Request.FormValue("password")
	registerValue.Ip = c.Request.FormValue("ip")
	err := repository.UserCreate(registerValue)
	fmt.Print(err)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "添加成功",
	})

}
