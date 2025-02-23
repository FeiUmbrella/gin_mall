package middleware

import (
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

// JWT 中间件，先对用户通过http传来的token进行解密验证，来确定封装在token中的用户信息。
// 如果解密验证不成功，不进行后续处理直接返回
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeOut
			}
		}

		if code != e.Success {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
