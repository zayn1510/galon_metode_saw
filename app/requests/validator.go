package requests

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
)

var (
	validate   *validator.Validate
	translator ut.Translator
)

func init() {
	// Inisialisasi translator
	eng := en.New()
	uni := ut.New(eng, eng)

	trans, _ := uni.GetTranslator("en")
	validate = validator.New()

	enTranslations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		if name == "-" {
			return ""
		}
		return name
	})

	translator = trans
}

func Validate(data any) (error, map[string]string) {
	err := validate.Struct(data)
	if err == nil {
		return nil, nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Translate(translator)
	}

	return err, errors
}
