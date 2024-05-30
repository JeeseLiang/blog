package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
)

type VaildError struct {
	Key     string
	Message string
}

type VaildErrors []*VaildError

func (e *VaildError) Error() string {
	return e.Message
}

func (e VaildErrors) Errors() []string {
	ret := []string{}
	for _, v := range e {
		ret = append(ret, v.Message)
	}
	return ret
}

func (e VaildErrors) Error() string {
	return strings.Join(e.Errors(), ",")
}

func BindAndVaild(c *gin.Context, v interface{}) (bool, VaildErrors) {
	var errs VaildErrors
	err := c.ShouldBind(v)
	if err != nil {
		trans, _ := c.Value("trans").(ut.Translator)

		verr, ok := err.(v10.ValidationErrors)

		if !ok {
			return false, errs
		}

		for key, msg := range verr.Translate(trans) {
			errs = append(errs, &VaildError{
				Key:     key,
				Message: msg,
			})
		}
		return false, errs
	}
	return true, nil
}
