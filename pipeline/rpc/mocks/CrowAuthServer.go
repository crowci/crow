// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	proto "github.com/crowci/crow/v3/pipeline/rpc/proto"
	mock "github.com/stretchr/testify/mock"
)

// CrowAuthServer is an autogenerated mock type for the CrowAuthServer type
type CrowAuthServer struct {
	mock.Mock
}

// Auth provides a mock function with given fields: _a0, _a1
func (_m *CrowAuthServer) Auth(_a0 context.Context, _a1 *proto.AuthRequest) (*proto.AuthResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Auth")
	}

	var r0 *proto.AuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.AuthRequest) (*proto.AuthResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.AuthRequest) *proto.AuthResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.AuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.AuthRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedCrowAuthServer provides a mock function with given fields:
func (_m *CrowAuthServer) mustEmbedUnimplementedCrowAuthServer() {
	_m.Called()
}

// NewCrowAuthServer creates a new instance of CrowAuthServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCrowAuthServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *CrowAuthServer {
	mock := &CrowAuthServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
