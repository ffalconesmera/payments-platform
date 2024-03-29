// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ffalconesmera/payments-platform/payments/repository (interfaces: PaymentRepository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	model "github.com/ffalconesmera/payments-platform/payments/model"
	gomock "github.com/golang/mock/gomock"
)

// MockPaymentRepository is a mock of PaymentRepository interface.
type MockPaymentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentRepositoryMockRecorder
}

// MockPaymentRepositoryMockRecorder is the mock recorder for MockPaymentRepository.
type MockPaymentRepositoryMockRecorder struct {
	mock *MockPaymentRepository
}

// NewMockPaymentRepository creates a new mock instance.
func NewMockPaymentRepository(ctrl *gomock.Controller) *MockPaymentRepository {
	mock := &MockPaymentRepository{ctrl: ctrl}
	mock.recorder = &MockPaymentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentRepository) EXPECT() *MockPaymentRepositoryMockRecorder {
	return m.recorder
}

// CreatePayment mocks base method.
func (m *MockPaymentRepository) CreatePayment(arg0 context.Context, arg1 *model.PayTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockPaymentRepositoryMockRecorder) CreatePayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockPaymentRepository)(nil).CreatePayment), arg0, arg1)
}

// FindPaymentByCode mocks base method.
func (m *MockPaymentRepository) FindPaymentByCode(arg0 context.Context, arg1 string) (*model.PayTransaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPaymentByCode", arg0, arg1)
	ret0, _ := ret[0].(*model.PayTransaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindPaymentByCode indicates an expected call of FindPaymentByCode.
func (mr *MockPaymentRepositoryMockRecorder) FindPaymentByCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPaymentByCode", reflect.TypeOf((*MockPaymentRepository)(nil).FindPaymentByCode), arg0, arg1)
}

// SavePayment mocks base method.
func (m *MockPaymentRepository) SavePayment(arg0 context.Context, arg1 *model.PayTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePayment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePayment indicates an expected call of SavePayment.
func (mr *MockPaymentRepositoryMockRecorder) SavePayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePayment", reflect.TypeOf((*MockPaymentRepository)(nil).SavePayment), arg0, arg1)
}
