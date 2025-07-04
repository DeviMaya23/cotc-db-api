// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"context"
	"lizobly/cotc-db-api/pkg/domain"

	mock "github.com/stretchr/testify/mock"
)

// NewMockUserService creates a new instance of MockUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserService {
	mock := &MockUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockUserService is an autogenerated mock type for the UserService type
type MockUserService struct {
	mock.Mock
}

type MockUserService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserService) EXPECT() *MockUserService_Expecter {
	return &MockUserService_Expecter{mock: &_m.Mock}
}

// Login provides a mock function for the type MockUserService
func (_mock *MockUserService) Login(ctx context.Context, req domain.LoginRequest) (domain.LoginResponse, error) {
	ret := _mock.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 domain.LoginResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, domain.LoginRequest) (domain.LoginResponse, error)); ok {
		return returnFunc(ctx, req)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, domain.LoginRequest) domain.LoginResponse); ok {
		r0 = returnFunc(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.LoginResponse)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, domain.LoginRequest) error); ok {
		r1 = returnFunc(ctx, req)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockUserService_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type MockUserService_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - req domain.LoginRequest
func (_e *MockUserService_Expecter) Login(ctx interface{}, req interface{}) *MockUserService_Login_Call {
	return &MockUserService_Login_Call{Call: _e.mock.On("Login", ctx, req)}
}

func (_c *MockUserService_Login_Call) Run(run func(ctx context.Context, req domain.LoginRequest)) *MockUserService_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 domain.LoginRequest
		if args[1] != nil {
			arg1 = args[1].(domain.LoginRequest)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockUserService_Login_Call) Return(res domain.LoginResponse, err error) *MockUserService_Login_Call {
	_c.Call.Return(res, err)
	return _c
}

func (_c *MockUserService_Login_Call) RunAndReturn(run func(ctx context.Context, req domain.LoginRequest) (domain.LoginResponse, error)) *MockUserService_Login_Call {
	_c.Call.Return(run)
	return _c
}
