package utils

import "regexp"

func IsValidPhoneNumber(phonenumber string) bool {
	//使用正则表达式
	reg := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return reg.MatchString(phonenumber)
}
