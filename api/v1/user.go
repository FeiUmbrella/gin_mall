package v1

import (
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controler
func UserRegister(c *gin.Context) {
	var userRegister service.UserService // 请求参数
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
