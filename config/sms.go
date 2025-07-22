package config

// import (
// 	"fmt"
// 	"gochat/dao/redis"
// 	"gochat/service"

// 	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
// 	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
// 	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
// )

// // InitSmsService 初始化短信服务
// func InitSmsService(redisDao *redis.SmsDao) (service.SmsService, error) {
// 	cfg := GlobalConfig.Sms

// 	// 校验必要配置
// 	if !cfg.Enable {
// 		return nil, fmt.Errorf("短信服务未启用")
// 	}

// 	if cfg.AccessKeyID == "" || cfg.AccessKeySecret == "" {
// 		return nil, fmt.Errorf("短信服务配置缺失: AccessKeyID/AccessKeySecret")
// 	}

// 	if cfg.SignName == "" || cfg.TemplateCode == "" {
// 		return nil, fmt.Errorf("短信服务配置缺失: SignName/TemplateCode")
// 	}

// 	// 设置默认区域
// 	if cfg.RegionID == "" {
// 		cfg.RegionID = "cn-hangzhou"
// 	}

// 	// 创建阿里云客户端
// 	clientConfig := sdk.NewConfig()
// 	credential := credentials.NewAccessKeyCredential(cfg.AccessKeyID, cfg.AccessKeySecret)
// 	client, err := dysmsapi.NewClientWithOptions(cfg.RegionID, clientConfig, credential)
// 	if err != nil {
// 		return nil, fmt.Errorf("创建阿里云短信客户端失败: %v", err)
// 	}

// 	// 创建短信服务实例
// 	return service.NewAliSmsService(
// 		redisDao,
// 		client,
// 		cfg.SignName,
// 		cfg.TemplateCode,
// 	), nil
// }
