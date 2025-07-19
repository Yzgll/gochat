package router

import (
	"gochat/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	//初始化路由
	r := gin.Default()
	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandler())

	UserRouterInit(r, db)
	return r

}
