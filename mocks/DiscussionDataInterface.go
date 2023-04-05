// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	discussions "lapakUmkm/features/discussions"

	mock "github.com/stretchr/testify/mock"
)

// DiscussionDataInterface is an autogenerated mock type for the DiscussionDataInterface type
type DiscussionDataInterface struct {
	mock.Mock
}

// Destroy provides a mock function with given fields: id
func (_m *DiscussionDataInterface) Destroy(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: discussionEntity, id
func (_m *DiscussionDataInterface) Edit(discussionEntity discussions.DiscussionEntity, id uint) error {
	ret := _m.Called(discussionEntity, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(discussions.DiscussionEntity, uint) error); ok {
		r0 = rf(discussionEntity, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields: userId
func (_m *DiscussionDataInterface) SelectAll(userId uint) ([]discussions.DiscussionEntity, error) {
	ret := _m.Called(userId)

	var r0 []discussions.DiscussionEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]discussions.DiscussionEntity, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uint) []discussions.DiscussionEntity); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]discussions.DiscussionEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: id
func (_m *DiscussionDataInterface) SelectById(id uint) (discussions.DiscussionEntity, error) {
	ret := _m.Called(id)

	var r0 discussions.DiscussionEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (discussions.DiscussionEntity, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) discussions.DiscussionEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(discussions.DiscussionEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDiscussionByProductId provides a mock function with given fields: productId
func (_m *DiscussionDataInterface) SelectDiscussionByProductId(productId uint) ([]discussions.DiscussionEntity, error) {
	ret := _m.Called(productId)

	var r0 []discussions.DiscussionEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]discussions.DiscussionEntity, error)); ok {
		return rf(productId)
	}
	if rf, ok := ret.Get(0).(func(uint) []discussions.DiscussionEntity); ok {
		r0 = rf(productId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]discussions.DiscussionEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: discussionEntity
func (_m *DiscussionDataInterface) Store(discussionEntity discussions.DiscussionEntity) (uint, error) {
	ret := _m.Called(discussionEntity)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(discussions.DiscussionEntity) (uint, error)); ok {
		return rf(discussionEntity)
	}
	if rf, ok := ret.Get(0).(func(discussions.DiscussionEntity) uint); ok {
		r0 = rf(discussionEntity)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(discussions.DiscussionEntity) error); ok {
		r1 = rf(discussionEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDiscussionDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDiscussionDataInterface creates a new instance of DiscussionDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDiscussionDataInterface(t mockConstructorTestingTNewDiscussionDataInterface) *DiscussionDataInterface {
	mock := &DiscussionDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}