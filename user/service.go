package user

import (
	"context"
	"lizobly/cotc-db/pkg/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func (s UserService) Login(ctx context.Context, req domain.LoginRequest) (res domain.LoginResponse, err error) {

	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		err = domain.ErrUserNotFound
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		err = domain.ErrInvalidPassword
		return
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	jwtTimeoutStr := os.Getenv("JWT_TIMEOUT")
	jwtTimeout, _ := time.ParseDuration(jwtTimeoutStr)

	exp := time.Now().Add(jwtTimeout)
	claims := domain.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return
	}

	res.Username = req.Username
	res.Token = t

	return
}
