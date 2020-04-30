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
	if err := c.ShouldBind(&registerValue); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}

	registerValue.Username = c.Request.FormValue("username")
	registerValue.Password = c.Request.FormValue("password")
	registerValue.Ip = c.Request.FormValue("ip")

	if registerValue.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "请输入账号",
		})
		return

	}
	err := repository.UserRead(registerValue)
	// err := repository.UserCreate(registerValue)
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
