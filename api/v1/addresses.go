package v1

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListAddress 展示所有地址-Controller
func ListAddress(c *gin.Context) {
	var listAddress service.AddressService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listAddress); err == nil {
		res := listAddress.List(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ShowAddress Api Error", err)
	}
}

// CreateAddress 创建地址-Controller
func CreateAddress(c *gin.Context) {
	var createAddress service.AddressService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createAddress); err == nil {
		res := createAddress.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateAddress Api Error", err)
	}
}

// DeleteAddress 删除某一地址-Controller
func DeleteAddress(c *gin.Context) {
	var deleteAddressService service.AddressService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteAddressService); err == nil {
		res := deleteAddressService.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteAddress Api Error", err)
	}
}

// GetAddress 得到某一地址-Controller
func GetAddress(c *gin.Context) {
	var getAddressService service.AddressService
	if err := c.ShouldBind(&getAddressService); err == nil {
		res := getAddressService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("GetAddress Api Error", err)
	}
}

// UpdateAddress 更新某一地址-Controller
func UpdateAddress(c *gin.Context) {
	var updateAddressService service.AddressService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateAddressService); err == nil {
		res := updateAddressService.Update(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("UpdateAddress Api Error", err)
	}
}
