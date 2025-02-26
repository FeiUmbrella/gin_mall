package v1

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOrder 创建订单-Controller
func CreateOrder(c *gin.Context) {
	var createOrder service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createOrder); err == nil {
		res := createOrder.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateOrder Api Error", err)
	}
}

// DeleteOrder 删除某一订单-Controller
func DeleteOrder(c *gin.Context) {
	var deleteOrderService service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteOrder Api Error", err)
	}
}

// ListOrder 显示所有订单-Controller
func ListOrder(c *gin.Context) {
	var listOrderService service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listOrderService); err == nil {
		res := listOrderService.List(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListOrder Api Error", err)
	}
}

// ShowOrder 显示某一订单详细内容-Controller
func ShowOrder(c *gin.Context) {
	var showOrderService service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ShowOrder Api Error", err)
	}
}
