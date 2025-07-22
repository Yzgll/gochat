package utils

import (
	"math/rand"
	"regexp"
	"time"
	"unicode"
)

func IsValidPhoneNumber(phonenumber string) bool {
	//使用正则表达式
	reg := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return reg.MatchString(phonenumber)
}

// 验证邮箱格式
func IsValidEmail(email string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return reg.MatchString(email)
}

// 验证密码强度
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false
	specialChars := map[rune]bool{
		'@': true, '$': true, '!': true, '%': true,
		'*': true, '?': true, '&': true,
	}
	for _, r := range password {
		switch {
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsDigit(r):
			hasDigit = true
		case specialChars[r]:
			hasSpecial = true
		}
	}
	return hasLower && hasUpper && hasDigit && hasSpecial
}

// 验证昵称格式
func IsValidNickname(nickname string) bool {
	// 2 - 16 个字符，允许中文、字母、数字、下划线
	// \p{Han} 匹配汉字，a-zA-Z0-9_ 匹配字母、数字、下划线
	reg := regexp.MustCompile(`^[\p{Han}a-zA-Z0-9_]{2,16}$`)
	return reg.MatchString(nickname)
}

// 验证出生日期格式
func IsValidBirthdate(birthdate string) bool {
	_, err := time.Parse("2006-01-02", birthdate)
	return err == nil
}

// 初始化随机数种子（确保每次程序启动时生成不同的随机序列）
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateSimpleGgnumber 生成简单的 10 位随机数（第一位不为 0）
func GenerateSimpleGgnumber() uint64 {
	// 第一位为 1-9
	firstDigit := rand.Intn(9) + 1
	// 后 9 位为 0-9 的随机数
	remainingDigits := rand.Int63n(1e9)

	// 组合成 10 位数
	return uint64(firstDigit)*1e9 + uint64(remainingDigits)
}
