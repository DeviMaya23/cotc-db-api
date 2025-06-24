package rest

import (
	"net/http"

	pkgValidator "lizobly/cotc-db-api/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"

	"github.com/labstack/echo/v4"
)

type Controller struct {
}

type StandardAPIResponse struct {
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Errors   interface{} `json:"errors"`
	Metadata interface{} `json:"metadata"`
}

type ValidationErrorFields struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (c Controller) Ok(ctx echo.Context, message string, data, metadata interface{}) error {

	return ctx.JSON(http.StatusOK, StandardAPIResponse{
		Message:  message,
		Data:     data,
		Metadata: metadata,
	})
}

func (c Controller) ResponseError(ctx echo.Context, httpStatus int, message string, errorData interface{}) error {

	return ctx.JSON(httpStatus, StandardAPIResponse{
		Message: message,
		Errors:  errorData,
	})
}

func (c Controller) ResponseErrorValidation(ctx echo.Context, err error) error {

	// TODO : non go validator error
	// _, ok := err.(*echo.HTTPError)
	// if !ok {
	// 	report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	var errMsg []ValidationErrorFields
	validate := ctx.Get("validator").(*pkgValidator.CustomValidator)
	language := ctx.Request().Header.Get("Accept-Language")
	translator, _ := validate.Translator.FindTranslator(language)

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, e := range castedObject {
			errMsg = append(errMsg, ValidationErrorFields{
				Field:   strcase.ToSnake(e.Field()),
				Message: e.Translate(translator),
			})
		}
	}

	return ctx.JSON(http.StatusBadRequest, StandardAPIResponse{
		Message: "error validation",
		Errors:  errMsg,
	})
}
