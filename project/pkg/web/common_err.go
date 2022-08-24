// 自定义错误

package web

import (
	"fmt"
	"net/http"
)

// 业务常用错误
var (
	ErrUnknown           = NewError("UnKnow", "未知错误")
	BadRequest           = NewError("BadRequest", "请求参数有误")
	ErrDB                = NewError("DBErr", "数据库发生错误")
	ErrUnauthorizedToken = NewError("UnauthorizedToken", "TOKEN 验证失败")
	ErrJSON              = NewError("UnmarshalErr", "JSON 编解码出错")
	ErrSystem            = NewError("SystemException", "系统异常")
)

// Error ...
type Error struct {
	reason  string   // 错误原因
	msg     string   // 错误信息，用户可读
	details []string // 错误扩展，开发可读
}

var codes = make(map[string]string, 8)

// NewError 创建自定义错误
func NewError(reason, msg string) *Error {
	if _, ok := codes[reason]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", reason))
	}
	codes[reason] = msg
	return &Error{reason: reason, msg: msg}
}

// Code ..
func (e *Error) Reason() string {
	return e.reason
}

// Message ..
func (e *Error) Message() string {
	return e.msg
}

// Details 错误
func (e *Error) Details() []string {
	return e.details
}

// WithDetails 错误详情
func (e *Error) WithDetails(args ...string) *Error {
	newErr := *e
	newErr.details = make([]string, 0, len(args))
	newErr.details = append(newErr.details, args...)
	return &newErr
}

// HTTPCode http status code
// 权限相关错误 401
// 程序错误 500
// 其它错误 400
func (e *Error) HTTPCode() int {
	switch e.reason {
	case "":
		return http.StatusOK
	case ErrUnauthorizedToken.reason:
		return http.StatusUnauthorized
	case ErrSystem.reason:
		return http.StatusInternalServerError
	}
	return http.StatusBadRequest
}
