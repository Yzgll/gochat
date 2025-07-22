package config

import (
	"context"
	"fmt"
	"gochat/utils"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	var err error
	rdsconfig := GlobalConfig.Redis
	addr := fmt.Sprintf("%s:%d", rdsconfig.Host, rdsconfig.Port)
	rdsclient := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     rdsconfig.Password,
		DB:           rdsconfig.DB,
		PoolSize:     rdsconfig.PoolSize,
		MinIdleConns: rdsconfig.MinIdleConns,
		MaxIdleConns: rdsconfig.MaxIdleConns,
	})
	//测试链接
	ctx := context.Background()
	if _, err = rdsclient.Ping(ctx).Result(); err != nil {
		utils.Log.Error("Redis数据库连接失败!")

	}
	utils.Log.Infof("Redis数据库连接成功！ ")
	return rdsclient
}
