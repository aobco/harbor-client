// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	icon "github.com/aobco/harbor-client/v5/apiv2/internal/api/client/icon"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockIconClientService is an autogenerated mock type for the ClientService type
type MockIconClientService struct {
	mock.Mock
}

// GetIcon provides a mock function with given fields: params, authInfo
func (_m *MockIconClientService) GetIcon(params *icon.GetIconParams, authInfo runtime.ClientAuthInfoWriter) (*icon.GetIconOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *icon.GetIconOK
	if rf, ok := ret.Get(0).(func(*icon.GetIconParams, runtime.ClientAuthInfoWriter) *icon.GetIconOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*icon.GetIconOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*icon.GetIconParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockIconClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
