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

		authed := v1.Group("/")      // 需要登录保护,将路由封装在v1下
		authed.Use(middleware.JWT()) // 使用JWT中间件
		{
			// 用户信息修改
			authed.PUT("user", api.UserUpdate)
			// 上传头像
			authed.POST("avatar", api.UploadAvatar)
		}
	}
	return r
}
