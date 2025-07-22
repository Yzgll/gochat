package redis

import (
	"context"
	"fmt"

	"time"

	"github.com/redis/go-redis/v9"
)

type SmsDao struct {
	Client *redis.Client
}

func NewSmsDao(client *redis.Client) *SmsDao {
	return &SmsDao{Client: client}
}

// SetCode 存储验证码到Redis
func (sd *SmsDao) SetCode(ctx context.Context, phone, code string, expire time.Duration) error {
	key := fmt.Sprintf("smscode:%s", phone)
	return sd.Client.Set(ctx, key, code, expire).Err()
}

// GetCode 从Redis获取验证码
func (sd *SmsDao) GetCode(ctx context.Context, phone string) (string, error) {
	key := fmt.Sprintf("smscode:%s", phone)
	return sd.Client.Get(ctx, key).Result()
}

// DelCode 验证成功后删除验证码
func (sd *SmsDao) DelCode(ctx context.Context, phone string) error {
	key := fmt.Sprintf("sms:code:%s", phone)
	return sd.Client.Del(ctx, key).Err()
}
