package v1

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCarts 创建购物车-Controller
func CreateCarts(c *gin.Context) {
	var createCarts service.CartService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCarts); err == nil {
		res := createCarts.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateCarts Api Error", err)
	}
}

// DeleteCarts 删除某一购物车-Controller
func DeleteCarts(c *gin.Context) {
	var deleteCartsService service.CartService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCartsService); err == nil {
		res := deleteCartsService.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteCarts Api Error", err)
	}
}

// ListCarts 显示购物车所有内容-Controller
func ListCarts(c *gin.Context) {
	var listCartsService service.CartService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listCartsService); err == nil {
		res := listCartsService.List(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("GetCarts Api Error", err)
	}
}

// UpdateCarts 更新某一购物车-Controller
func UpdateCarts(c *gin.Context) {
	var updateCartsService service.CartService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateCartsService); err == nil {
		res := updateCartsService.Update(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("UpdateCarts Api Error", err)
	}
}
