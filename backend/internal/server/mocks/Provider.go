// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

type Provider_Expecter struct {
	mock *mock.Mock
}

func (_m *Provider) EXPECT() *Provider_Expecter {
	return &Provider_Expecter{mock: &_m.Mock}
}

// Validate provides a mock function with given fields: provider, token
func (_m *Provider) Validate(provider string, token string) (string, bool, error) {
	ret := _m.Called(provider, token)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 string
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (string, bool, error)); ok {
		return rf(provider, token)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(provider, token)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) bool); ok {
		r1 = rf(provider, token)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(provider, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Provider_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type Provider_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - provider string
//   - token string
func (_e *Provider_Expecter) Validate(provider interface{}, token interface{}) *Provider_Validate_Call {
	return &Provider_Validate_Call{Call: _e.mock.On("Validate", provider, token)}
}

func (_c *Provider_Validate_Call) Run(run func(provider string, token string)) *Provider_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Provider_Validate_Call) Return(_a0 string, _a1 bool, _a2 error) *Provider_Validate_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Provider_Validate_Call) RunAndReturn(run func(string, string) (string, bool, error)) *Provider_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewProvider creates a new instance of Provider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *Provider {
	mock := &Provider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
