package validate

import (
	"log"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

type BindValidate struct{}

var Bv *BindValidate

// BindValidate
// i18n
func (bv *BindValidate) BindValidate() (ut.Translator, *validator.Validate) {
	validate := validator.New()
	e := en.New()
	uniTrans := ut.New(e, e, zh.New())
	translator, _ := uniTrans.GetTranslator("zh")
	err := zh2.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		log.Printf("i18n 国际化失败: %v", err)
	}
	return translator, validate
}
