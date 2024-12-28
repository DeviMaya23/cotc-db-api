package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type CustomValidator struct {
	Validator  *validator.Validate
	Translator *ut.UniversalTranslator
}

func NewValidator() *CustomValidator {

	en := en.New()
	uni := ut.New(en, en)

	english, ok := uni.GetTranslator("en")
	if !ok {
		fmt.Println("failed get en translator")
	}
	newValidator := validator.New()
	en_translations.RegisterDefaultTranslations(newValidator, english)

	return &CustomValidator{
		Validator:  newValidator,
		Translator: uni,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		return err
	}
	return nil
}
