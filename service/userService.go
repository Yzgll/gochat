package service

import (
	"errors"
	"gochat/dto"
	"gochat/model"
	"gochat/utils"

	"gorm.io/gorm"
)

// UserService 处理用户相关业务逻辑
type UserService struct {
	db *gorm.DB // 依赖注入的 db 实例

}

// 工厂函数创建实例
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// 处理用户注册
func (us *UserService) Register(req dto.RegisterRequest) (*model.User, error) {
	//检查各种字段是否符合要求
	//1.检查手机号
	if !utils.IsValidPhoneNumber(req.Phone) {
		return nil, errors.New("手机号格式错误，请输入11位手机号！")
	}
	//检查邮箱格式

}


