package mysql

import (
	"fmt"
	"gochat/model"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// ExistByPhone 检查手机号是否已注册
func (ud *UserDao) ExistByPhone(phone string) bool {
	var count int64
	ud.db.Model(&model.User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// ExistByEmail 检查邮箱是否已注册
func (ud *UserDao) ExistByEmail(email string) bool {
	var count int64
	ud.db.Model(&model.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// 创建用户
func (ud *UserDao) CreateUser(user *model.User) error {
	return ud.db.Create(user).Error
}

// 根据gg号查询用户信息
func (ud *UserDao) GetUserByGGnumber(ggnumber uint64) (user *model.User, err error) {
	user = &model.User{}
	result := ud.db.Where("ggnumber=?", ggnumber).First(user)
	err = result.Error
	//未找到
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	// 处理其他错误
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}
	return user, nil //查询成功
}

// 根据手机号查询
func (ud *UserDao) GetUserByPhone(ggnumber uint64) (user *model.User, err error) {}

// 根据邮箱查询
func (ud *UserDao) GetUserByEmail(ggnumber uint64) (user *model.User, err error) {}
