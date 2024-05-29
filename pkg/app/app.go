package app

import (
	"blog/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct { // 返回的响应信息
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"totol_rows"`
}

func NewResponse(c *gin.Context) *Response {
	return &Response{
		Ctx: c,
	}
}

func (r *Response) ToResponse(data interface{}) interface{} {
	if data == nil {
		data = gin.H{}
	}
	return data
}

func (r *Response) ToResponseList(dataList interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": dataList,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	details := err.Details()
	r.Ctx.JSON(err.StatusCode(), gin.H{
		"code":    err.Code(),
		"msg":     err.Msg(),
		"details": details,
	})
}
