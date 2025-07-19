package main

import (
	"gochat/config"
	"gochat/model"
	"gochat/router"
	"gochat/utils"
	"strconv"
)

func main() {
	//加载配置文件
	config.InitConfig()
	//初始化数据库
	db := config.InitMysql()
	model.InitModel(db)
	//初始化日志
	utils.InitLogger()

	// 注册路由
	r := router.InitRouter(db)
	port := strconv.Itoa(config.GlobalConfig.System.Port)
	utils.Log.Infof("服务启动在端口: %s", port)
	r.Run(":" + port)
}
