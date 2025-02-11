package rest

import (
	"errors"
	"lizobly/cotc-db/pkg/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserService interface {
	Login(ctx echo.Context, req domain.LoginRequest) (res domain.LoginResponse, err error)
}

type UserHandler struct {
	Controller
	Service UserService
}

func NewUserHandler(e *echo.Group, svc UserService) {
	handler := &UserHandler{
		Service: svc,
	}

	e.POST("/login", handler.Login)
}

func (h *UserHandler) Login(ctx echo.Context) error {

	var request domain.LoginRequest

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = ctx.Validate(&request)
	if err != nil {
		return h.ResponseErrorValidation(ctx, err)
	}

	res, err := h.Service.Login(ctx, request)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidPassword):
			return h.ResponseError(ctx, http.StatusBadRequest, "error", err.Error())
		case errors.Is(err, domain.ErrUserNotFound):
			return h.ResponseError(ctx, http.StatusBadRequest, "error", err.Error())
		default:
			return h.ResponseError(ctx, http.StatusInternalServerError, "error", err.Error())
		}
	}

	return h.Ok(ctx, "success", res, nil)
}
