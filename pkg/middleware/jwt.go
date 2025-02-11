package middleware

import (
	"lizobly/cotc-db/pkg/domain"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewJWTMiddleware() echo.MiddlewareFunc {

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	cfg := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.JWTClaims)
		},
		SigningKey: []byte(jwtSecretKey),
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/api/v1/login"
		},
	}

	return echojwt.WithConfig(cfg)
}
