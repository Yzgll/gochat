package service

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"gochat/dao/redis"
// 	"gochat/utils"
// 	"strings"
// 	"time"

// 	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
// )

// type SmsService interface {
// 	//发送验证码
// 	Send(phone string) error

// 	//验证验证码
// 	Verify(phone, code string) bool
// }

// // AliSmsService 阿里云短信服务实现
// type AliSmsService struct {
// 	redisDao *redis.SmsDao // Redis 操作
// 	client   *dysmsapi.Client
// 	signName string // 短信签名
// 	template string // 模板ID
// }

// // 工厂函数创建短信服务实例
// func NewAliSmsService(redisDao *redis.SmsDao, aliclient *dysmsapi.Client, signName, template string) *AliSmsService {
// 	return &AliSmsService{
// 		redisDao: redisDao,
// 		client:   aliclient,
// 		signName: signName,
// 		template: template,
// 	}
// }

// // 发送验证码函数
// func (as *AliSmsService) Send(phone string) error {
// 	//1.检查手机号格式
// 	if !utils.IsValidPhoneNumber(phone) {
// 		utils.Log.Warnf("手机号格式错误: %s", phone)
// 		return errors.New("手机号格式错误，请输入正确的手机号！")
// 	}
// 	//2.限制频率一分钟发送一次
// 	ctx := context.Background()
// 	if err := as.checkRateLimit(ctx, phone); err != nil {
// 		utils.Log.Warnf("手机号 %s 发送频率超限: %v", phone, err)
// 		return errors.New("发送过于频繁，请1分钟后再试")
// 	}
// 	//3.生成六位数验证码
// 	code := utils.GenerateCode(6)
// 	//4.调用阿里云短信api
// 	if err := as.callAliyunAPI(phone, code); err != nil {
// 		utils.Log.Errorf("手机号 %s 短信发送失败: %v", phone, err)
// 		return errors.New("短信发送失败，请稍后重试")
// 	}
// 	//5.验证码存入redis，有效期五分钟
// 	if err := as.redisDao.SetCode(ctx, phone, code, 5*time.Minute); err != nil {
// 		utils.Log.Errorf("手机号 %s 验证码缓存失败: %v", phone, err)

// 	}
// 	// 6. 记录发送日志（审计用）
// 	utils.Log.Infof("手机号 %s 短信发送成功（验证码已缓存）", phone)
// 	return nil

// }

// // 验证验证码
// func (as *AliSmsService) Verify(phone, code string) bool {
// 	// 1. 基础参数校验
// 	if phone == "" || code == "" || len(code) != 6 {
// 		utils.Log.Warnf("手机号 %s 验证码格式错误: %s", phone, code)
// 		return false
// 	}
// 	//2.从redis缓存读取验证码
// 	ctx := context.Background()
// 	cachecode, err := as.redisDao.GetCode(ctx, phone)
// 	if err != nil {
// 		// 区分"未找到"和"Redis错误"
// 		if strings.Contains(err.Error(), "redis: nil") {
// 			utils.Log.Infof("手机号 %s 验证码已过期", phone)
// 		} else {
// 			utils.Log.Errorf("手机号 %s 验证码查询失败: %v", phone, err)

// 		}
// 		return false
// 	}
// 	//3.对比验证码
// 	if cachecode != code {
// 		utils.Log.Warnf("手机号 %s 验证码不匹配（用户输入: %s，缓存: %s）", phone, code, cachecode)
// 		return false
// 	}

// 	// 4. 验证成功后立即删除验证码（防止重复使用）
// 	if err := as.redisDao.DelCode(ctx, phone); err != nil {
// 		utils.Log.Warnf("手机号 %s 验证码删除失败: %v", phone, err)
// 		// 不影响验证结果，但需记录告警
// 	}
// 	utils.Log.Infof("手机号 %s 验证码验证通过", phone)
// 	return true
// }

// // 调用阿里云短信API（私有方法，封装第三方调用）
// func (as *AliSmsService) callAliyunAPI(phone, code string) error {
// 	request := dysmsapi.CreateSendSmsRequest()
// 	request.Scheme = "https" // 强制HTTPS加密传输
// 	request.PhoneNumbers = phone
// 	request.SignName = as.signName
// 	request.TemplateCode = as.template
// 	request.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, code)

// 	// 调用阿里云SDK发送短信
// 	response, err := as.client.SendSms(request)
// 	if err != nil {
// 		return fmt.Errorf("阿里云API调用失败: %v", err)
// 	}

// 	// 处理阿里云返回的错误码（详细错误处理）
// 	if response.Code != "OK" {
// 		return fmt.Errorf("阿里云返回错误: 代码=%s, 信息=%s", response.Code, response.Message)
// 	}

// 	return nil
// }

// // 频率限制（私有方法，防止刷短信）
// func (as *AliSmsService) checkRateLimit(ctx context.Context, phone string) error {
// 	// 用Redis的SETNX实现1分钟内只能发送1次
// 	key := fmt.Sprintf("sms:ratelimit:%s", phone)
// 	// 设置键值为"1"，过期时间1分钟，NX确保只有不存在时才设置成功
// 	ok, err := as.redisDao.Client.SetNX(ctx, key, "1", time.Minute).Result()
// 	if err != nil {
// 		return fmt.Errorf("频率限制Redis操作失败: %v", err)
// 	}
// 	if !ok {
// 		return errors.New("1分钟内只能发送1条短信")
// 	}
// 	return nil
// }
