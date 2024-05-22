package errcode

// 定义公共错误码
var (
	Success                        = NewError(0, "成功")
	NotFound                       = NewError(10001, "找不到页面")
	InvalidParams                  = NewError(10002, "参数错误")
	ServerError                    = NewError(10003, "服务器错误")
	TooManyRequest                 = NewError(10004, "请求过多")
	UnauthoerizedAuthNotExist      = NewError(10005, "鉴权失败，找不到对应的appkey")
	UnauthoerizedTokenError        = NewError(10006, "鉴权失败，Token错误")
	UnauthoerizedTokenTimeout      = NewError(10007, "鉴权失败，Token超时")
	UnauthoerizedTokenGeneralError = NewError(10008, "鉴权失败，Token生成失败")
)
