// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	models "github.com/VolunteerOne/volunteer-one-app/backend/models"
	mock "github.com/stretchr/testify/mock"
)

// FriendService is an autogenerated mock type for the FriendService type
type FriendService struct {
	mock.Mock
}

// AcceptFriend provides a mock function with given fields: friend
func (_m *FriendService) AcceptFriend(friend models.Friend) (models.Friend, error) {
	ret := _m.Called(friend)

	var r0 models.Friend
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Friend) (models.Friend, error)); ok {
		return rf(friend)
	}
	if rf, ok := ret.Get(0).(func(models.Friend) models.Friend); ok {
		r0 = rf(friend)
	} else {
		r0 = ret.Get(0).(models.Friend)
	}

	if rf, ok := ret.Get(1).(func(models.Friend) error); ok {
		r1 = rf(friend)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFriend provides a mock function with given fields: friend
func (_m *FriendService) CreateFriend(friend models.Friend) (models.Friend, error) {
	ret := _m.Called(friend)

	var r0 models.Friend
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Friend) (models.Friend, error)); ok {
		return rf(friend)
	}
	if rf, ok := ret.Get(0).(func(models.Friend) models.Friend); ok {
		r0 = rf(friend)
	} else {
		r0 = ret.Get(0).(models.Friend)
	}

	if rf, ok := ret.Get(1).(func(models.Friend) error); ok {
		r1 = rf(friend)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFriends provides a mock function with given fields:
func (_m *FriendService) GetFriends() ([]models.Friend, error) {
	ret := _m.Called()

	var r0 []models.Friend
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Friend, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Friend); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Friend)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OneFriend provides a mock function with given fields: id
func (_m *FriendService) OneFriend(id string) (models.Friend, error) {
	ret := _m.Called(id)

	var r0 models.Friend
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Friend, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Friend); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Friend)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RejectFriend provides a mock function with given fields: friend
func (_m *FriendService) RejectFriend(friend models.Friend) error {
	ret := _m.Called(friend)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Friend) error); ok {
		r0 = rf(friend)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFriendService interface {
	mock.TestingT
	Cleanup(func())
}

// NewFriendService creates a new instance of FriendService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFriendService(t mockConstructorTestingTNewFriendService) *FriendService {
	mock := &FriendService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
