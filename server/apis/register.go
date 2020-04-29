package apis

import (
	"net/http"
	models "server/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerDb models.Register
	registerDb.Username = c.Request.FormValue("username")
	registerDb.Password = c.Request.FormValue("password")
	registerDb.Ip = c.Request.FormValue("ip")
	err := registerDb.Insert()

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
