package middleware

import (
	"gochat/utils"
	"time"

	"github.com/gin-gonic/gin"
)

//日志中间件

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		//记录开始时间
		start_time := time.Now()

		c.Next()
		//记录请求的日志
		duration := time.Since(start_time)
		statusCode := c.Writer.Status()

		utils.Log.Infof("[ACCESS] %s %s | Status: %d | Duration: %v | Client IP: %s",
			c.Request.Method, c.Request.URL.Path, statusCode, duration, c.ClientIP())
	}
}
