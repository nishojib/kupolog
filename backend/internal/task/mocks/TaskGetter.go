// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	task "github.com/nishojib/ffxivdailies/internal/task"
	mock "github.com/stretchr/testify/mock"
)

// TaskGetter is an autogenerated mock type for the TaskGetter type
type TaskGetter struct {
	mock.Mock
}

type TaskGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *TaskGetter) EXPECT() *TaskGetter_Expecter {
	return &TaskGetter_Expecter{mock: &_m.Mock}
}

// GetTasksForUser provides a mock function with given fields: ctx, userID
func (_m *TaskGetter) GetTasksForUser(ctx context.Context, userID string) ([]task.Task, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetTasksForUser")
	}

	var r0 []task.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]task.Task, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []task.Task); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]task.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskGetter_GetTasksForUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTasksForUser'
type TaskGetter_GetTasksForUser_Call struct {
	*mock.Call
}

// GetTasksForUser is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *TaskGetter_Expecter) GetTasksForUser(ctx interface{}, userID interface{}) *TaskGetter_GetTasksForUser_Call {
	return &TaskGetter_GetTasksForUser_Call{Call: _e.mock.On("GetTasksForUser", ctx, userID)}
}

func (_c *TaskGetter_GetTasksForUser_Call) Run(run func(ctx context.Context, userID string)) *TaskGetter_GetTasksForUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TaskGetter_GetTasksForUser_Call) Return(_a0 []task.Task, _a1 error) *TaskGetter_GetTasksForUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TaskGetter_GetTasksForUser_Call) RunAndReturn(run func(context.Context, string) ([]task.Task, error)) *TaskGetter_GetTasksForUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewTaskGetter creates a new instance of TaskGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskGetter {
	mock := &TaskGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
