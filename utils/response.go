package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 通用的返回
func Response(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"result":  data,
	})
}

// Success 成功的请求
func Success(ctx *gin.Context, data interface{}) {
	Response(ctx, 0, "请求成功", data)
	return
}

// Fail 失败的请求
func Fail(ctx *gin.Context, message string) {
	Response(ctx, 1, message, nil)
}