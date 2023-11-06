// Code generated by mockery v2.30.1. DO NOT EDIT.

package usecase_test

import (
	command "github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryUpdateProfileMock is an autogenerated mock type for the RepositoryUpdateProfile type
type RepositoryUpdateProfileMock struct {
	mock.Mock
}

// UpdateClassesDoneInUserProfile provides a mock function with given fields: cmd
func (_m *RepositoryUpdateProfileMock) UpdateClassesDoneInUserProfile(cmd command.UpdateClassesDone) error {
	ret := _m.Called(cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(command.UpdateClassesDone) error); ok {
		r0 = rf(cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositoryUpdateProfileMock creates a new instance of RepositoryUpdateProfileMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryUpdateProfileMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryUpdateProfileMock {
	mock := &RepositoryUpdateProfileMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}