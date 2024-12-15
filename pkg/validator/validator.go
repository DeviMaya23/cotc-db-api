package validator

import (
	"fmt"
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func NewValidator() *CustomValidator {

	en := en.New()
	uni := ut.New(en, en)

	english, ok := uni.GetTranslator("en")
	if !ok {
		fmt.Println("failed get en translator")
	}
	newValidator := validator.New()
	en_translations.RegisterDefaultTranslations(newValidator, english)

	// newValidator.RegisterTranslation("required", english, func(ut ut.Translator) error {
	// 	return ut.Add("required", "{0} must have a value!", true)
	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("required", fe.Field())

	// 	return t
	// })

	return &CustomValidator{
		validator: newValidator,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// TODO :
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
