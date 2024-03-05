// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ffalconesmera/payments-platform/payments/config (interfaces: Config)

// Package mock_config is a generated GoMock package.
package mock_config

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConfig is a mock of Config interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig.
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// GetDatabaseHost mocks base method.
func (m *MockConfig) GetDatabaseHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabaseHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDatabaseHost indicates an expected call of GetDatabaseHost.
func (mr *MockConfigMockRecorder) GetDatabaseHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabaseHost", reflect.TypeOf((*MockConfig)(nil).GetDatabaseHost))
}

// GetDatabaseName mocks base method.
func (m *MockConfig) GetDatabaseName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabaseName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDatabaseName indicates an expected call of GetDatabaseName.
func (mr *MockConfigMockRecorder) GetDatabaseName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabaseName", reflect.TypeOf((*MockConfig)(nil).GetDatabaseName))
}

// GetDatabasePassword mocks base method.
func (m *MockConfig) GetDatabasePassword() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabasePassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDatabasePassword indicates an expected call of GetDatabasePassword.
func (mr *MockConfigMockRecorder) GetDatabasePassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabasePassword", reflect.TypeOf((*MockConfig)(nil).GetDatabasePassword))
}

// GetDatabasePort mocks base method.
func (m *MockConfig) GetDatabasePort() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabasePort")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDatabasePort indicates an expected call of GetDatabasePort.
func (mr *MockConfigMockRecorder) GetDatabasePort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabasePort", reflect.TypeOf((*MockConfig)(nil).GetDatabasePort))
}

// GetDatabaseUser mocks base method.
func (m *MockConfig) GetDatabaseUser() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabaseUser")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDatabaseUser indicates an expected call of GetDatabaseUser.
func (mr *MockConfigMockRecorder) GetDatabaseUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabaseUser", reflect.TypeOf((*MockConfig)(nil).GetDatabaseUser))
}

// GetJWTExpiration mocks base method.
func (m *MockConfig) GetJWTExpiration() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJWTExpiration")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetJWTExpiration indicates an expected call of GetJWTExpiration.
func (mr *MockConfigMockRecorder) GetJWTExpiration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJWTExpiration", reflect.TypeOf((*MockConfig)(nil).GetJWTExpiration))
}

// GetJWTSecretKey mocks base method.
func (m *MockConfig) GetJWTSecretKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJWTSecretKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetJWTSecretKey indicates an expected call of GetJWTSecretKey.
func (mr *MockConfigMockRecorder) GetJWTSecretKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJWTSecretKey", reflect.TypeOf((*MockConfig)(nil).GetJWTSecretKey))
}

// InitConfig mocks base method.
func (m *MockConfig) InitConfig() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitConfig")
}

// InitConfig indicates an expected call of InitConfig.
func (mr *MockConfigMockRecorder) InitConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitConfig", reflect.TypeOf((*MockConfig)(nil).InitConfig))
}
