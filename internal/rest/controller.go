package rest

import (
	"net/http"

	pkgValidator "lizobly/cotc-db/pkg/validator"

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

	// _, ok := err.(*echo.HTTPError)
	// if !ok {
	// 	report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	var validationMessages []ValidationErrorFields
	ctxValidator := ctx.Get("validator").(*pkgValidator.CustomValidator)

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, e := range castedObject {
			validationMessages = append(validationMessages, ValidationErrorFields{
				Field:   strcase.ToSnake(e.Field()),
				Message: e.Translate(ctxValidator.Translator.GetFallback()),
			})
		}
	}

	return ctx.JSON(http.StatusInternalServerError, StandardAPIResponse{
		Message: "error validation",
		Errors:  validationMessages,
	})
}
