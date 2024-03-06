// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ffalconesmera/payments-platform/payments/externals/repository (interfaces: MerchantRepository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	dto "github.com/ffalconesmera/payments-platform/payments/externals/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockMerchantRepository is a mock of MerchantRepository interface.
type MockMerchantRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMerchantRepositoryMockRecorder
}

// MockMerchantRepositoryMockRecorder is the mock recorder for MockMerchantRepository.
type MockMerchantRepositoryMockRecorder struct {
	mock *MockMerchantRepository
}

// NewMockMerchantRepository creates a new mock instance.
func NewMockMerchantRepository(ctrl *gomock.Controller) *MockMerchantRepository {
	mock := &MockMerchantRepository{ctrl: ctrl}
	mock.recorder = &MockMerchantRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMerchantRepository) EXPECT() *MockMerchantRepositoryMockRecorder {
	return m.recorder
}

// FindMerchantByCode mocks base method.
func (m *MockMerchantRepository) FindMerchantByCode(arg0 string) (*dto.Merchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMerchantByCode", arg0)
	ret0, _ := ret[0].(*dto.Merchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMerchantByCode indicates an expected call of FindMerchantByCode.
func (mr *MockMerchantRepositoryMockRecorder) FindMerchantByCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMerchantByCode", reflect.TypeOf((*MockMerchantRepository)(nil).FindMerchantByCode), arg0)
}