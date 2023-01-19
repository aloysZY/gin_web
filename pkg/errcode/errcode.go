// Package errcode 对错误码的封装
package errcode

import (
	"fmt"
	"net/http"
)

// Error 错误结构体
type Error struct {
	code    int      `json:"code"`    // 错误码
	msg     string   `json:"msg"`     // 错误消息
	details []string `json:"details"` // 详细信息
}

// 实现方法的好处是，创建出来的变量直接可以"."调用
var codes = map[int]string{}

// NewError 自定义错误K,V,返回错误结构体,这是在 new 错误信息的时候进行确认
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

// ErrorF 格式化输出单个错误码和错误信息
func (e *Error) ErrorF() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

// Code 返回单个错误码
func (e *Error) Code() int {
	return e.code
}

// Msg 返回单个错误MSG信息
func (e *Error) Msg() string {
	return e.msg
}

// MsgF 返回多个错误MSG信息，根据传入错误
func (e *Error) MsgF(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// Details 返回单个错误详情信息
func (e *Error) Details() []string {
	return e.details
}

// WithDetails 返回多个错误详情信息，可变参数
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

// StatusCode 根据错误码，匹配返回的状态码
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
