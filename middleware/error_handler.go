package middleware

import (
	"fmt"
	"gochat/errors"
	"gochat/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一处理错误中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("panic: %v", r)
				}
				utils.Log.Errorf("[PANIC] path=%s,err=%v", c.Request.URL.Path, err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "服务器内部错误，请稍后重试",
				})
				c.Abort()
			}
		}()
		c.Next()
		if len(c.Errors) == 0 {
			return
		}
		err := c.Errors.Last().Err
		if errors.IsAppError(err) {
			e := err.(errors.AppError)
			statusCode := c.Writer.Status()
			if statusCode == http.StatusOK {
				statusCode = e.HTTPStatus()
			}
			utils.Log.Errorf("发生自定义错误 | 路径: %s | 错误类型: %v | 错误码: %s | 提示: %s | 原始错误: %v", c.Request.URL.Path,
				e.Type(), e.Code(), e.Message(), e.OriginalError())
			c.JSON(statusCode, gin.H{
				"success": false,
				"message": e.Message(),
			})
		} else {
			utils.Log.Errorf("未知错误：%v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(), // 这里返回普通 error 的 message
			})
		}
	}
}
