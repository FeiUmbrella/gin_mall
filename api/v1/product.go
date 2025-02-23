package v1

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	// todo: gin框架下通过http传文件的相关函数
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var createProductService service.ProductService
	// 将http传来的参数放进定义的showMoney结构体中
	if err := c.ShouldBind(&createProductService); err == nil {
		// claims中存放着登录用户的ID，将其传入service层的Update函数中
		res := createProductService.Create(c.Request.Context(), claims.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateProduct Api Error: ", err)
	}
}

func ListProduct(c *gin.Context) {
	var listProductService service.ProductService
	if err := c.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListProduct Api Error: ", err)
	}
}

func SearchProduct(c *gin.Context) {
	var searchProductService service.ProductService
	if err := c.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("SearchProduct Api Error: ", err)
	}
}

// ShowProduct 获取特定id商品的详细信息
func ShowProduct(c *gin.Context) {
	var showProductService service.ProductService
	if err := c.ShouldBind(&showProductService); err == nil {
		res := showProductService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ShowProduct Api Error: ", err)
	}
}
