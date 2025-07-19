package dto

// 注册请求的结构体
type RegisterRequest struct {
	//基础信息
	Password  string `json:"password" binding:"required,min=8"` //binding强制要求字段存在且不为空
	Nickname  string `json:"nickname" binding:"required,min=2,max=16"`
	Birthdate string `json:"birthdate" binding:"required,len=10"` //yyyy-mm-dd
	Phone     string `json:"phone" binding:"required,len=11,number"`
	Email     string `json:"email" binding:"required,email"`

	//验证码
	SmsCode string `json:"sms_code" binding:"required,len=6,number"` // 短信验证码
}

//发送验证码请求结构体
type SendSmsRequest struct {
	Phone string `json:"phone" binding:required,len=11,number`
}

//验证验证码结构体
type VerifySmsRequest struct {
	Phone   string `json:"phone" binding:"required,len=11,number"`
	SmsCode string `json:"smscode" binding:"required,len=6,number"`
}
