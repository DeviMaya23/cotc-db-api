package rest

import (
	"encoding/json"
	"lizobly/cotc-db-api/internal/rest/mocks"
	"lizobly/cotc-db-api/pkg/domain"
	"lizobly/cotc-db-api/pkg/helpers"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserHandlerSuite struct {
	suite.Suite

	e           *echo.Echo
	userService *mocks.MockUserService
	handler     *UserHandler
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuite))
}

func (s *UserHandlerSuite) SetupSuite() {

	s.e = echo.New()
	s.userService = new(mocks.MockUserService)
	s.handler = NewUserHandler(s.e.Group(""), s.userService)

}

func (s *UserHandlerSuite) TestUserHandler_NewHandler() {

	got := NewUserHandler(s.e.Group(""), s.userService)
	assert.Equal(s.T(), got, s.handler)

}

func (s *UserHandlerSuite) TestTravellerHandler_Login() {

	type args struct {
		requestBody interface{}
	}
	type want struct {
		loginResponse domain.LoginResponse
		responseBody  interface{}
		statusCode    int
	}

	login := domain.LoginRequest{
		Username: "user",
		Password: "pw",
	}
	resp := domain.LoginResponse{
		Username: "isla",
		Token:    "1234",
	}

	tests := []struct {
		name       string
		args       args
		want       want
		beforeTest func(ctx echo.Context, param args, want want)
	}{
		{
			name: "success get traveller",
			args: args{login},
			want: want{
				loginResponse: resp,
				responseBody: StandardAPIResponse{
					Message: "success",
					Data:    resp,
				},
				statusCode: http.StatusOK,
			},
			beforeTest: func(ctx echo.Context, param args, want want) {
				s.userService.On("Login", mock.Anything, param.requestBody).Return(want.loginResponse, nil).Once()

			},
		},
		{
			name: "failed bind",
			args: args{`asdf`},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "failed validation",
			args: args{domain.LoginRequest{}},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "failed ErrInvalidPassword",
			args: args{login},
			want: want{
				loginResponse: resp,
				responseBody: StandardAPIResponse{
					Message: "error",
					Errors:  domain.ErrInvalidPassword.Error(),
				},
				statusCode: http.StatusBadRequest,
			},
			beforeTest: func(ctx echo.Context, param args, want want) {
				s.userService.On("Login", mock.Anything, param.requestBody).Return(want.loginResponse, domain.ErrInvalidPassword).Once()

			},
		},
		{
			name: "failed ErrUserNotFound",
			args: args{login},
			want: want{
				loginResponse: resp,
				responseBody: StandardAPIResponse{
					Message: "error",
					Errors:  domain.ErrUserNotFound.Error(),
				},
				statusCode: http.StatusBadRequest,
			},
			beforeTest: func(ctx echo.Context, param args, want want) {
				s.userService.On("Login", mock.Anything, param.requestBody).Return(want.loginResponse, domain.ErrUserNotFound).Once()

			},
		},
		{
			name: "failed internal server error",
			args: args{login},
			want: want{
				loginResponse: resp,
				responseBody: StandardAPIResponse{
					Message: "error",
					Errors:  gorm.ErrCheckConstraintViolated.Error(),
				},
				statusCode: http.StatusInternalServerError,
			},
			beforeTest: func(ctx echo.Context, param args, want want) {
				s.userService.On("Login", mock.Anything, param.requestBody).Return(want.loginResponse, gorm.ErrCheckConstraintViolated).Once()

			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {

			rec, ctx := helpers.GetHTTPTestRecorder(s.T(), http.MethodPost, "/login", tt.args.requestBody, nil, nil)

			if tt.beforeTest != nil {
				tt.beforeTest(ctx, tt.args, tt.want)
			}

			err := s.handler.Login(ctx)
			assert.Nil(s.T(), err)
			assert.Equal(s.T(), tt.want.statusCode, ctx.Response().Status)

			if tt.want.responseBody != nil {

				wantRespBytes, err := json.Marshal(tt.want.responseBody)
				assert.NoError(s.T(), err)

				assert.Equal(s.T(), string(wantRespBytes), strings.TrimSpace(rec.Body.String()))

			}

		})
	}

}
