package main

import (
	"gochat/config"
	"gochat/dao/mysql"
	"gochat/router"
	"gochat/utils"
	"strconv"
)

func main() {
	//初始化日志
	utils.InitLogger()
	//加载配置文件
	config.InitConfig()
	//初始化数据库
	mysqldb := config.InitMysql()
	//redis
	redisdb := config.InitRedis()

	userdao := mysql.NewUserDAO(mysqldb)
	// 注册路由
	r := router.InitRouter(redisdb, userdao)
	port := strconv.Itoa(config.GlobalConfig.System.Port)
	utils.Log.Infof("服务启动在端口: %s", port)
	r.Run(":" + port)
}
