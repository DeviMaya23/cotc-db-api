package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type CustomValidator struct {
	Validator  *validator.Validate
	Translator *ut.UniversalTranslator
}

func NewValidator() *CustomValidator {

	newValidator := validator.New()

	en := en.New()
	id := id.New()

	uni := ut.New(en, en, id)

	english, ok := uni.GetTranslator("en")
	if !ok {
		fmt.Println("failed get en translator")
	}
	en_translations.RegisterDefaultTranslations(newValidator, english)

	indonesian, ok := uni.GetTranslator("id")
	if !ok {
		fmt.Println("failed get id translator")
	}
	id_translations.RegisterDefaultTranslations(newValidator, indonesian)

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
