package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int
	msg     string
	details []string
}

var Err = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := Err[code]; ok {
		panic(fmt.Sprintf("错误码{ %d }已经被占用！", code))
	}
	Err[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details []string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码 : %d\n错误信息 : %s\n", e.code, e.msg)
}

func (e *Error) StatusCode() int {
	switch e.code {
	case Success.code:
		return http.StatusOK // 200
	case NotFound.code:
		return http.StatusNotFound // 404
	case InvalidParams.code:
		return http.StatusBadRequest // 400
	case UnauthoerizedAuthNotExist.code:
		return http.StatusUnauthorized // 401
	case UnauthoerizedTokenError.code:
		return http.StatusUnauthorized // 401
	case UnauthoerizedTokenTimeout.code:
		return http.StatusUnauthorized // 401
	case UnauthoerizedTokenGeneralError.code:
		return http.StatusUnauthorized // 401
	case TooManyRequest.code:
		return http.StatusTooManyRequests // 429
	}
	return http.StatusInternalServerError // 500
}
