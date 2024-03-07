// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/manuelbilbao/cli/v28/ignite/services/plugin/grpc/v1"
)

// PluginClientAPI is an autogenerated mock type for the ClientAPI type
type PluginClientAPI struct {
	mock.Mock
}

type PluginClientAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *PluginClientAPI) EXPECT() *PluginClientAPI_Expecter {
	return &PluginClientAPI_Expecter{mock: &_m.Mock}
}

// GetChainInfo provides a mock function with given fields: _a0
func (_m *PluginClientAPI) GetChainInfo(_a0 context.Context) (*v1.ChainInfo, error) {
	ret := _m.Called(_a0)

	var r0 *v1.ChainInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*v1.ChainInfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *v1.ChainInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ChainInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PluginClientAPI_GetChainInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChainInfo'
type PluginClientAPI_GetChainInfo_Call struct {
	*mock.Call
}

// GetChainInfo is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *PluginClientAPI_Expecter) GetChainInfo(_a0 interface{}) *PluginClientAPI_GetChainInfo_Call {
	return &PluginClientAPI_GetChainInfo_Call{Call: _e.mock.On("GetChainInfo", _a0)}
}

func (_c *PluginClientAPI_GetChainInfo_Call) Run(run func(_a0 context.Context)) *PluginClientAPI_GetChainInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *PluginClientAPI_GetChainInfo_Call) Return(_a0 *v1.ChainInfo, _a1 error) *PluginClientAPI_GetChainInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PluginClientAPI_GetChainInfo_Call) RunAndReturn(run func(context.Context) (*v1.ChainInfo, error)) *PluginClientAPI_GetChainInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewPluginClientAPI creates a new instance of PluginClientAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPluginClientAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *PluginClientAPI {
	mock := &PluginClientAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
