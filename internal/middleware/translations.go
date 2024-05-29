package middleware

// 翻译英文信息中间件

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	tr_en "github.com/go-playground/validator/v10/translations/en"
	tr_zh "github.com/go-playground/validator/v10/translations/zh"
)

func Translation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := ctx.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "en":
				_ = tr_en.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = tr_zh.RegisterDefaultTranslations(v, trans)
				break
			}
			ctx.Set("trans", trans)
		}
		ctx.Next()
	}
}
