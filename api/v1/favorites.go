package v1

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListFavorites 展示收藏内容Controller
func ListFavorites(c *gin.Context) {
	var listFavorites service.FavoritesService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listFavorites); err == nil {
		res := listFavorites.List(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ShowFavorites Api Error", err)
	}
}

// CreateFavorites 创建收藏内容Controller
func CreateFavorites(c *gin.Context) {
	var createFavorites service.FavoritesService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createFavorites); err == nil {
		res := createFavorites.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("CreateFavorites Api Error", err)
	}
}

// DeleteFavorites 删除收藏内容Controller
func DeleteFavorites(c *gin.Context) {
	var deleteFavoritesService service.FavoritesService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteFavoritesService); err == nil {
		res := deleteFavoritesService.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("DeleteFavorites Api Error", err)
	}
}
