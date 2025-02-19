package v1

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controler层

func UserRegister(c *gin.Context) {
	var userRegister service.UserService // 请求参数
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService // 请求参数
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService // 请求参数
	// 从http的header中取出token，验证token
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err == nil {
		// claims中存放着登录用户的ID，将其传入service层的Update函数中
		res := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService // 请求参数
	// 从http的header中取出token，验证token
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatar); err == nil {
		// claims中存放着登录用户的ID，将其传入service层的Update函数中
		res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService
	// 从http的header中取出token，解析token中的参数
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	// 将http传来的参数放进定义的sendEmail结构体中
	if err := c.ShouldBind(&sendEmail); err == nil {
		// claims中存放着登录用户的ID，将其传入service层的Update函数中
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService
	// 将http传来的参数放进定义的validEmail结构体中
	if err := c.ShouldBind(&validEmail); err == nil {
		// claims中存放着登录用户的ID，将其传入service层的Update函数中
		res := validEmail.Send(c.Request.Context(), c.GetHeader("Authorization"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
