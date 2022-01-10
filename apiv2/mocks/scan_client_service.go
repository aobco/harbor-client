// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	scan "github.com/aobco/goharbor-client/v5/apiv2/internal/api/client/scan"
)

// MockScanClientService is an autogenerated mock type for the ClientService type
type MockScanClientService struct {
	mock.Mock
}

// GetReportLog provides a mock function with given fields: params, authInfo
func (_m *MockScanClientService) GetReportLog(params *scan.GetReportLogParams, authInfo runtime.ClientAuthInfoWriter) (*scan.GetReportLogOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan.GetReportLogOK
	if rf, ok := ret.Get(0).(func(*scan.GetReportLogParams, runtime.ClientAuthInfoWriter) *scan.GetReportLogOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan.GetReportLogOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan.GetReportLogParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanArtifact provides a mock function with given fields: params, authInfo
func (_m *MockScanClientService) ScanArtifact(params *scan.ScanArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*scan.ScanArtifactAccepted, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan.ScanArtifactAccepted
	if rf, ok := ret.Get(0).(func(*scan.ScanArtifactParams, runtime.ClientAuthInfoWriter) *scan.ScanArtifactAccepted); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan.ScanArtifactAccepted)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan.ScanArtifactParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockScanClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// StopScanArtifact provides a mock function with given fields: params, authInfo
func (_m *MockScanClientService) StopScanArtifact(params *scan.StopScanArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*scan.StopScanArtifactAccepted, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan.StopScanArtifactAccepted
	if rf, ok := ret.Get(0).(func(*scan.StopScanArtifactParams, runtime.ClientAuthInfoWriter) *scan.StopScanArtifactAccepted); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan.StopScanArtifactAccepted)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan.StopScanArtifactParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
