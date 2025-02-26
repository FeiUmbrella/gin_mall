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

		// 轮播图
		v1.GET("carousels", api.ListCarousel)

		// 商品列表
		v1.GET("products", api.ListProduct)
		// 获取特定id商品的详细信息
		v1.GET("products/:id", api.ShowProduct)
		// 获取商品的图片信息
		v1.GET("imgs/:id", api.ListProductImg)
		// 获取商品分类
		v1.GET("categories", api.ListCategories)

		authed := v1.Group("/")      // 需要登录保护,将路由封装在v1下
		authed.Use(middleware.JWT()) // 使用JWT中间件
		{
			// 用户信息修改
			authed.PUT("user", api.UserUpdate)
			// 上传头像
			authed.POST("avatar", api.UploadAvatar)
			// 发送邮件
			authed.POST("user/sending-email", api.SendEmail)
			// 验证向用户邮箱发送的包含token的链接
			authed.POST("user/valid-email", api.ValidEmail)

			// 显示用户金额
			authed.POST("money", api.ShowMoney)

			// 创建商品--用户都可以创建商品. 类似咸鱼
			authed.POST("product", api.CreateProduct)
			// 搜索商品
			authed.POST("products", api.SearchProduct)

			// 显示收藏夹
			authed.GET("favorites", api.ListFavorites)
			// 创建收藏
			authed.POST("favorites", api.CreateFavorites)
			// 删除收藏
			authed.DELETE("favorites/:id", api.DeleteFavorites)

			// 上传地址
			authed.POST("addresses/", api.CreateAddress)
			// 获取某一地址
			authed.GET("addresses/:id", api.GetAddress)
			// 显示所有地址
			authed.GET("addresses/", api.ListAddress)
			// 更新某一地址
			authed.PUT("addresses/:id", api.UpdateAddress)
			// 删除地址
			authed.DELETE("addresses/:id", api.DeleteAddress)

			// 新建购物车
			authed.POST("carts", api.CreateCarts)
			// 获取购物车内所有信息
			authed.GET("carts", api.ListCarts)
			// 更新某一购物车
			authed.PUT("carts/:id", api.UpdateCarts)
			// 删除购物车
			authed.DELETE("carts/:id", api.DeleteCarts)

			// 新建订单
			authed.POST("order", api.CreateOrder)
			// 获取订单
			authed.GET("order", api.ListOrder)
			// 获取订单详细信息
			authed.GET("order/:id", api.ShowOrder)
			// 删除订单
			authed.DELETE("order/:id", api.DeleteOrder)
		}
	}
	return r
}
