package router

import (
	"gochat/controller"
	"gochat/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouterInit(r *gin.Engine, db *gorm.DB) {
	// 1. 初始化 service 和 controller
	userSvc := service.NewUserService(db) //注入数据库db
	userCtrl := controller.NewUserController(userSvc)
	userRouter := r.Group("/api/user")
	{
		userRouter.POST("/register", userCtrl.Register)
	}
}
