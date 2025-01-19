package middleware

import (
	"lizobly/cotc-db/pkg/domain"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewJWTMiddleware() echo.MiddlewareFunc {

	cfg := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.JWTClaims)
		},
		SigningKey: []byte("catnipsforisla"),
	}

	return echojwt.WithConfig(cfg)
}
