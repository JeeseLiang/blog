package app

import (
	"blog/global"
	"blog/pkg"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := pkg.Str(c.Query("page")).OnlyInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	page_size := pkg.Str(c.Query("page_size")).OnlyInt()
	if page_size >= global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	if page_size <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	return page_size
}

func GetPageOffset(page, page_size int) int {
	off := (page - 1) * page_size
	if off <= 0 {
		return 0
	}
	return off
}
