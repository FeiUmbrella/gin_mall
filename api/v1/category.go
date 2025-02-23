package v1

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategories(c *gin.Context) {
	var listCategory service.CategoryService // 请求参数
	if err := c.ShouldBind(&listCategory); err == nil {
		res := listCategory.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListCategories Api Error: ", err)
	}
}
