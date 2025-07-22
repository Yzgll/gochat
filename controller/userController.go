package controller

import (
	"gochat/dto"
	"gochat/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController 处理用户相关 HTTP 请求
type UserController struct {
	UserService *service.UserService
}

// NewUserController 创建 UserController 实例
func NewUserController(svc *service.UserService) *UserController {
	return &UserController{UserService: svc}
}

// 发送验证码
// func (uc *UserController) SendSMS(c *gin.Context) {
// 	var req dto.SendSmsRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"code":    400,
// 			"message": "手机号格式错误: " + err.Error(),
// 		})
// 		return
// 	}

// 	// 调用 Service 层发送短信
// 	err := uc.UserService.SendSMS(req.Phone)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"code":    400,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code":    200,
// 		"message": "验证码已发送，有效期5分钟",
// 	})
// }

// 注册接口
// func (uc *UserController) Register(c *gin.Context) {
// var req dto.RegisterRequest
// if err := c.ShouldBindJSON(&req); err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"code":    400,
// 		"message": "参数格式错误: " + err.Error(),
// 	})
// 	return
// }

// // 1. 先验证验证码（调用短信服务）
// if !uc.UserService.VerifySMS(req.Phone, req.SmsCode) {
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"code":    400,
// 		"message": "验证码错误或已过期",
// 	})
// 	return
// }

// // 2. 调用用户服务完成注册
// user, err := uc.UserService.Register(req)
// if err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"code":    400,
// 		"message": err.Error(),
// 	})
// 	return
// }

// // 注册成功
// c.JSON(http.StatusOK, gin.H{
// 	"code":    200,
// 	"message": "注册成功",
// 	"data": gin.H{
// 		"ggnumber": user.Ggnumber,
// 	},
// })
// }

func (uc *UserController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := uc.UserService.Register(req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "注册成功",
		"data": map[string]interface{}{
			"id":        user.ID,
			"ggnumber":  user.Ggnumber,
			"nickname":  user.Nickname,
			"avatarurl": user.Avatarurl,
		},
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

}
