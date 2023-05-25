package utils

import (
	validate2 "github.com/echoframe/pkg/validate"
	"github.com/go-playground/validator/v10"
)

var (
	msg string
)

// ValidateParam
func ValidateParam(param interface{}) string {
	translator, validate := validate2.Bv.BindValidate()
	err := validate.Struct(param)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, er := range errs {
			msg = er.Translate(translator)
		}
		return msg
	}
	return ""
}
