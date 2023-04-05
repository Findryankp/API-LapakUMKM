// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	chats "lapakUmkm/features/chats"

	mock "github.com/stretchr/testify/mock"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// AllMessageToMe provides a mock function with given fields: myId, userId
func (_m *ServiceInterface) AllMessageToMe(myId uint, userId uint) ([]chats.ChatEntity, error) {
	ret := _m.Called(myId, userId)

	var r0 []chats.ChatEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) ([]chats.ChatEntity, error)); ok {
		return rf(myId, userId)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) []chats.ChatEntity); ok {
		r0 = rf(myId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chats.ChatEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(myId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: chatEntity
func (_m *ServiceInterface) Create(chatEntity chats.ChatEntity) (chats.ChatEntity, error) {
	ret := _m.Called(chatEntity)

	var r0 chats.ChatEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(chats.ChatEntity) (chats.ChatEntity, error)); ok {
		return rf(chatEntity)
	}
	if rf, ok := ret.Get(0).(func(chats.ChatEntity) chats.ChatEntity); ok {
		r0 = rf(chatEntity)
	} else {
		r0 = ret.Get(0).(chats.ChatEntity)
	}

	if rf, ok := ret.Get(1).(func(chats.ChatEntity) error); ok {
		r1 = rf(chatEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *ServiceInterface) GetById(id uint) (chats.ChatEntity, error) {
	ret := _m.Called(id)

	var r0 chats.ChatEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (chats.ChatEntity, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) chats.ChatEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(chats.ChatEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRoomId provides a mock function with given fields: roomId
func (_m *ServiceInterface) GetByRoomId(roomId string) ([]chats.ChatEntity, error) {
	ret := _m.Called(roomId)

	var r0 []chats.ChatEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]chats.ChatEntity, error)); ok {
		return rf(roomId)
	}
	if rf, ok := ret.Get(0).(func(string) []chats.ChatEntity); ok {
		r0 = rf(roomId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chats.ChatEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(roomId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSenderUser provides a mock function with given fields: myId, userId
func (_m *ServiceInterface) GetSenderUser(myId uint, userId uint) ([]chats.ChatEntity, error) {
	ret := _m.Called(myId, userId)

	var r0 []chats.ChatEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) ([]chats.ChatEntity, error)); ok {
		return rf(myId, userId)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) []chats.ChatEntity); ok {
		r0 = rf(myId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chats.ChatEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(myId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceInterface creates a new instance of ServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceInterface(t mockConstructorTestingTNewServiceInterface) *ServiceInterface {
	mock := &ServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}