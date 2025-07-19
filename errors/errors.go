package errors

import (
	"fmt"
	"net/http"
)

type ErrorType int

const (
	ErrorTypeValidation   ErrorType = iota + 1 // 参数验证错误 (HTTP 400)
	ErrorTypeConflict                          // 资源冲突 (HTTP 409)
	ErrorTypeUnauthorized                      // 未授权 (HTTP 401)
	ErrorTypeForbidden                         // 禁止访问 (HTTP 403)
	ErrorTypeNotFound                          // 资源不存在 (HTTP 404)
	ErrorTypeInternal                          // 内部服务器错误 (HTTP 500)
)

//统一错误接口

type AppError interface {
	error
	Code() string         //返回细粒度的错误
	Type() ErrorType      //返回错误类型
	HTTPStatus() int      //返回错误码
	Message() string      //返回错误信息
	OriginalError() error //返回底层错误信息
}

// 实现接口的具体类
type BaseError struct {
	Err       error
	ErrorType ErrorType
	UserMsg   string //展示给用户看的信息
	LogMsg    string
}

// 实现接口的方法
func (e *BaseError) Error() string {

	if e.Err != nil {
		return e.Err.Error()
	}
	return e.UserMsg
}

func (e *BaseError) Code() string {
	//返回更细的错误分类，比如409的手机或邮箱已经注册
	return fmt.Sprintf("ERR_%03d", e.ErrorType)
}
func (e *BaseError) Type() ErrorType {
	return e.ErrorType
}

func (e *BaseError) HTTPStatus() int {
	switch e.ErrorType {
	case ErrorTypeValidation:
		return http.StatusBadRequest
	case ErrorTypeConflict:
		return http.StatusConflict
	case ErrorTypeUnauthorized:
		return http.StatusUnauthorized
	case ErrorTypeForbidden:
		return http.StatusForbidden
	case ErrorTypeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

// 返回给用户看的错误信息
func (e *BaseError) Message() string {
	return e.UserMsg
}

// 返回底层错误信息
func (e *BaseError) OriginalError() error {
	return e.Err
}

// 工厂函数创建AppError实例
func NewAppError(errorType ErrorType, userMsg string, originalErr ...error) AppError {
	var err error
	if len(originalErr) > 0 {
		err = originalErr[0]
	}
	return &BaseError{
		Err:       err,
		ErrorType: errorType,
		UserMsg:   userMsg,
		LogMsg:    fmt.Sprintf("%s: %v", userMsg, originalErr),
	}
}

// 封装底层错误
func Wrap(originalerr error, errortype ErrorType, usermsg string) AppError {
	return NewAppError(errortype, usermsg, originalerr)
}

// 判断错误类型
func IsAppError(err error) bool {
	_, ok := err.(AppError)
	return ok
}
