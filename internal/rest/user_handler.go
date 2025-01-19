package rest

import (
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

	e.POST("login", handler.Login)
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
		return h.ResponseError(ctx, http.StatusBadRequest, "error get data", err.Error())
	}

	return h.Ok(ctx, "success", res, nil)
}
