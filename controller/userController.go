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

//注册接口

func (uc *UserController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	//从http请求中解析json数据，根据binding验证，并填充到req中
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err) //将错误添加到上下文，由中间件处理
		return
	}
	// 调用 service 层注册逻辑
	user, err := uc.UserService.Register(req)
	if err != nil {
		c.Error(err)
		return
	}

	// 注册成功返回 Ggnumber
	c.JSON(http.StatusOK, gin.H{
		"message":  "注册成功",
		"ggnumber": user.Ggnumber,
	})
}
