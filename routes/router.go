package routes

import (
	api "gin_mall/api/v1"
	"gin_mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo: gin
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static")) // 加载静态文件
	v1 := r.Group("/api/v1")                    // 路由组放在 /api/v1 下
	{
		// 先ping一下是否连通
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		// ===用户操作===
		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
	}
	return r
}
