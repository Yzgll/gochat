package service

import (
	"errors"
	"gochat/dao/mysql"
	"gochat/dto"
	"gochat/model"
	"gochat/utils"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

// UserService 处理用户相关业务逻辑
type UserService struct {
	userdao *mysql.UserDao // 依赖注入的 db 实例
	redisdb *redis.Client
}

// 工厂函数创建实例
func NewUserService(userdao *mysql.UserDao, redisdb *redis.Client) *UserService {
	return &UserService{userdao: userdao,
		redisdb: redisdb,
	}
}

// 发送验证码
// func (us *UserService) SendSMS(phone string) error {
// 	return us.smsService.Send(phone)
// }
// func (us *UserService) VerifySMS(phone, code string) bool {
// 	return us.smsService.Verify(phone, code)
// }

// 处理用户注册
func (us *UserService) Register(req dto.RegisterRequest) (*model.User, error) {
	//检查各种字段是否符合要求
	//1.检查手机号
	if !utils.IsValidPhoneNumber(req.Phone) {
		return nil, errors.New("手机号格式错误，请输入11位手机号！")
	}

	if !utils.IsValidEmail(req.Email) {
		return nil, errors.New("邮箱格式错误，请输入有效的邮箱地址")
	}
	if !utils.IsValidBirthdate(req.Birthdate) {
		return nil, errors.New("出生日期格式错误，应为YYYY-MM-DD")
	}
	if !utils.IsValidPassword(req.Password) {
		return nil, errors.New("密码强度不足，需包含大小写字母、数字和特殊字符")
	}
	if !utils.IsValidNickname(req.Nickname) {
		return nil, errors.New("昵称格式错误，长度应为2-16个字符")
	}
	if us.userdao.ExistByPhone(req.Phone) {
		return nil, errors.New("该手机号已注册")
	}
	if us.userdao.ExistByEmail(req.Email) {
		return nil, errors.New("该邮箱已注册")
	}
	// 3. 验证短信验证码（调用短信服务）
	// if !us.smsService.Verify(req.Phone, req.SmsCode) {
	// 	return nil, errors.New("验证码错误或已过期")
	// }
	// 4. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败，请重试")
	}

	// 7. 创建用户对象
	user := &model.User{
		Phone:     req.Phone,              // 来自注册请求
		Email:     req.Email,              // 来自注册请求
		Nickname:  req.Nickname,           // 来自注册请求
		Password:  string(hashedPassword), // 加密后的密码
		Birthdate: req.Birthdate,          // 来自注册请求（格式：yyyy-mm-dd）
	}
	user.Ggnumber = utils.GenerateSimpleGgnumber()
	// 8. 写入数据库（使用事务确保数据一致性）
	err = us.userdao.CreateUser(user)
	if err != nil {
		return nil, errors.New("用户创建失败，请重试" + err.Error())
	}

	// 9. 注册成功后，记录注册日志
	utils.Log.Infof("用户注册成功 [ID: %d, Ggnumber: %s]", user.ID, user.Ggnumber)

	return user, nil
}

func (us *UserService) Login(req dto.LoginRequest) (*model.User, error) {
	var user *model.User
	var err error
	//根据输入的账户不同查询用户信息
	if req.Ggnumber > 0 {
		//根据GG号查询用户
		user, err = us.userdao.GetUserByGGnumber(req.Ggnumber)
	}
	if req.Phone != "" {
		user, err = us.userdao.GetUserByGGnumber(req.Ggnumber)
	}
	if req.Email != "" {
		user, err = us.userdao.GetUserByGGnumber(req.Ggnumber)
	}
}
