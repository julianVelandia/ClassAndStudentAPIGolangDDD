// Code generated by mockery v2.30.1. DO NOT EDIT.

package read_test

import (
	domain "github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/domain"
	dto "github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/infrastructure/repository/json/dto"
	mock "github.com/stretchr/testify/mock"
)

// MapperMock is an autogenerated mock type for the Mapper type
type MapperMock struct {
	mock.Mock
}

// DTOClassToDomain provides a mock function with given fields: class
func (_m *MapperMock) DTOClassToDomain(class dto.Class) domain.Class {
	ret := _m.Called(class)

	var r0 domain.Class
	if rf, ok := ret.Get(0).(func(dto.Class) domain.Class); ok {
		r0 = rf(class)
	} else {
		r0 = ret.Get(0).(domain.Class)
	}

	return r0
}

// NewMapperMock creates a new instance of MapperMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMapperMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *MapperMock {
	mock := &MapperMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}