package utils

import (
	"math/rand"
	"strings"
	"time"
)

// 初始化随机数种子（确保每次运行生成不同的随机序列）
func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateCode(length int) string {
	if length <= 0 {
		return ""
	}

	// 定义数字字符集
	digits := "0123456789"
	// 预分配足够长度的字符串构建器
	var sb strings.Builder
	sb.Grow(length)

	// 循环生成随机数字
	for i := 0; i < length; i++ {
		// 从字符集中随机选一个字符
		sb.WriteByte(digits[rand.Intn(len(digits))])
	}

	return sb.String()
}
