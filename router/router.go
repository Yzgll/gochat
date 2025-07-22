package router

import (
	"gochat/dao/mysql"
	"gochat/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func InitRouter(redisdb *redis.Client, userdao *mysql.UserDao) *gin.Engine {
	//初始化路由
	r := gin.Default()
	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandler())
	r.Use(cors.Default()) // 这里用 gin-contrib/cors

	UserRouterInit(r, redisdb, userdao)
	return r

}
