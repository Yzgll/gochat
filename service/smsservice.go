package service

import (
	"errors"
	"gochat/dto"
	"gochat/utils"

	"github.com/go-redis/redis/v8"
)

type SmsService interface {
	//发送验证码
	Send(phone string) (string, error)

	//验证验证码
	Verify(phone, code string) (bool, error)
}

// 实现短信服务结构体
type AliSmsService struct {
	client *aliyunSmsClient
	redis  *redis.Client
}

// 工厂函数创建实例
func NewAliSmsService() *AliSmsService {
	return &AliSmsService{}
}

// 发送验证码函数
func (as *AliSmsService) Send(req dto.SendSmsRequest) error {
	//检查手机号格式
	if !utils.IsValidPhoneNumber(req.Phone) {
		return errors.New("手机号错误，请输入正确的手机号！")
	}

}

func (as *AliSmsService) Verify() {}
