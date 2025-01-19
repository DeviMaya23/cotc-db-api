package user

import (
	"context"
	"errors"
	"lizobly/cotc-db/pkg/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (result domain.User, err error)
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(u UserRepository) *UserService {
	return &UserService{
		userRepo: u,
	}
}

func (s UserService) Login(ctx echo.Context, req domain.LoginRequest) (res domain.LoginResponse, err error) {

	user, err := s.userRepo.GetByUsername(ctx.Request().Context(), req.Username)
	if err != nil {
		err = errors.New("failed get user info")
		return
	}

	// TODO : hash password
	if user.Password != req.Password {
		err = errors.New("invalid password")
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	t, err := token.SignedString([]byte("2catnipsforisla"))
	if err != nil {
		return
	}

	res.Username = req.Username
	res.Token = t

	return
}
