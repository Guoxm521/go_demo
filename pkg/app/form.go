package app

import (
	"example.com/m/v2/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	InitTranslate()
	if err := c.ShouldBind(form); err != nil {
		//todo 后面需要记录日志
		return http.StatusInternalServerError, e.ERROR
	}
	return http.StatusOK, e.SUCCESS
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func InitTranslate() {
	zh2 := zh2.New()
	uni = ut.New(zh2, zh2)
	trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	zh.RegisterDefaultTranslations(validate, trans)
}

func Translate(err error) map[string][]string {
	var result = make(map[string][]string)
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
}
