// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	types "github.com/getamis/alice/types"
)

// EchoMessage is an autogenerated mock type for the EchoMessage type
type EchoMessage struct {
	mock.Mock
}

// GetEchoMessage provides a mock function with given fields:
func (_m *EchoMessage) GetEchoMessage() types.Message {
	ret := _m.Called()

	var r0 types.Message
	if rf, ok := ret.Get(0).(func() types.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Message)
		}
	}

	return r0
}

// GetId provides a mock function with given fields:
func (_m *EchoMessage) GetId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetMessageType provides a mock function with given fields:
func (_m *EchoMessage) GetMessageType() types.MessageType {
	ret := _m.Called()

	var r0 types.MessageType
	if rf, ok := ret.Get(0).(func() types.MessageType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.MessageType)
	}

	return r0
}

// IsValid provides a mock function with given fields:
func (_m *EchoMessage) IsValid() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ProtoReflect provides a mock function with given fields:
func (_m *EchoMessage) ProtoReflect() protoreflect.Message {
	ret := _m.Called()

	var r0 protoreflect.Message
	if rf, ok := ret.Get(0).(func() protoreflect.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.Message)
		}
	}

	return r0
}

// NewEchoMessage creates a new instance of EchoMessage. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewEchoMessage(t testing.TB) *EchoMessage {
	mock := &EchoMessage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}