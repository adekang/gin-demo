package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data any, msg string) {

	ctx.JSON(httpStatus, gin.H{
		"code":    code,
		"result":  data,
		"msg":     msg,
		"success": code == 200,
	})

}

func Success(ctx *gin.Context, data any, msg string) {

	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data any, msg string) {

	Response(ctx, http.StatusOK, 400, data, msg)
}
