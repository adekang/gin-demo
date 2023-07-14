package middleware

import (
	"github/adekang/gin-demo/common"
	"github/adekang/gin-demo/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// 获取Authorization Header
		tokenStr := ctx.GetHeader("Authorization")

		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenStr = tokenStr[7:]

		token, claims, err := common.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 通过claim 获取userId

		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在"})
			ctx.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()

	}

}
