package dto

// 注册请求的结构体
type RegisterRequest struct {
	//基础信息
	Password  string `json:"password" binding:"required,min=8"` //binding强制要求字段存在且不为空
	Nickname  string `json:"nickname" binding:"required,min=2,max=16"`
	Birthdate string `json:"birthdate" binding:"required,len=10"` //yyyy-mm-dd
	Phone     string `json:"phone" binding:"required,len=11,number"`
	Email     string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	// 三种登录方式任选其一（通过binding标签的验证规则实现互斥）
	Ggnumber uint64 `json:"ggnumber"`                                // Ggnumber是数字类型，非必填
	Phone    string `json:"phone" binding:"omitempty,len=11,number"` // 可选，长度11位数字
	Email    string `json:"email" binding:"omitempty,email"`         // 可选，邮箱格式

	Password string `json:"password" binding:"required,min=8"` // 密码必填
}
