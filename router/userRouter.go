package router

import (
	"gochat/controller"
	"gochat/dao/mysql"
	"gochat/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func UserRouterInit(r *gin.Engine, redisdb *redis.Client, userdao *mysql.UserDao) {
	// 1. 初始化 service 和 controller
	userSvc := service.NewUserService(userdao, redisdb) //注入数据库db
	userCtrl := controller.NewUserController(userSvc)

	userRouter := r.Group("/api/user")
	{

		userRouter.POST("/register", userCtrl.Register)

	}
}
