package rest

import (
	"net/http"

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

	return ctx.JSON(http.StatusInternalServerError, StandardAPIResponse{
		Message: "error validation",
		Errors:  err,
	})
}
