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
				//处理panic
				err, ok := r.(error)
				//转换失败，不是error类型，手动创建
				if !ok {
					err = fmt.Errorf("panic: %v", r)
				}
				//记录错误日志
				utils.Log.Errorf("[PANIC] path=%s,err=%v", c.Request.URL.Path, err)
				//返回同意错误相应
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "ERR_INTERNAL",
					"msg":   "服务器内部错误，请稍后重试",
				})
				c.Abort()
			}
		}()
		c.Next() // 继续执行后续中间件和处理函数
		// 处理所有错误（记录日志）
		for _, e := range c.Errors {
			utils.Log.Errorf("[ERROR] path=%s, error=%v",
				c.Request.URL.Path,
				e.Error(),
			)
		}
		// 如果没有错误，直接返回
		if len(c.Errors) == 0 {
			return
		}
		err := c.Errors.Last().Err //获取最后一个错误
		//判断是不是自定义类型
		if errors.IsAppError(err) {
			e := err.(errors.AppError)

			// 使用已设置的状态码或AppError提供的状态码
			statusCode := c.Writer.Status()
			if statusCode == http.StatusOK {
				statusCode = e.HTTPStatus()
			}
			//记录日志
			utils.Log.Errorf("发生自定义错误 | 路径: %s | 错误类型: %v | 错误码: %s | 提示: %s | 原始错误: %v", c.Request.URL.Path,
				e.Type(), e.Code(), e.Message(), e.OriginalError())

			//返回相应给客户端
			c.JSON(e.HTTPStatus(), gin.H{
				"error_code": e.Code(),
				"message":    e.Message(),
			})
		} else {
			utils.Log.Errorf("未知错误：%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "ERR_UNKNOWN",
				"msg":   "服务器开小差了，稍后再试！",
			})
		}
	}
}
