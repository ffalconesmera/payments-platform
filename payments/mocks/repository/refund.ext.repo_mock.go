// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ffalconesmera/payments-platform/payments/repository (interfaces: RefundRepository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	model "github.com/ffalconesmera/payments-platform/payments/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRefundRepository is a mock of RefundRepository interface.
type MockRefundRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRefundRepositoryMockRecorder
}

// MockRefundRepositoryMockRecorder is the mock recorder for MockRefundRepository.
type MockRefundRepositoryMockRecorder struct {
	mock *MockRefundRepository
}

// NewMockRefundRepository creates a new mock instance.
func NewMockRefundRepository(ctrl *gomock.Controller) *MockRefundRepository {
	mock := &MockRefundRepository{ctrl: ctrl}
	mock.recorder = &MockRefundRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefundRepository) EXPECT() *MockRefundRepositoryMockRecorder {
	return m.recorder
}

// CreateRefund mocks base method.
func (m *MockRefundRepository) CreateRefund(arg0 context.Context, arg1 *model.PayRefund) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefund", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRefund indicates an expected call of CreateRefund.
func (mr *MockRefundRepositoryMockRecorder) CreateRefund(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefund", reflect.TypeOf((*MockRefundRepository)(nil).CreateRefund), arg0, arg1)
}

// FindRefundById mocks base method.
func (m *MockRefundRepository) FindRefundById(arg0 context.Context, arg1 string) (*model.PayRefund, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRefundById", arg0, arg1)
	ret0, _ := ret[0].(*model.PayRefund)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindRefundById indicates an expected call of FindRefundById.
func (mr *MockRefundRepositoryMockRecorder) FindRefundById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRefundById", reflect.TypeOf((*MockRefundRepository)(nil).FindRefundById), arg0, arg1)
}

// SaveRefund mocks base method.
func (m *MockRefundRepository) SaveRefund(arg0 context.Context, arg1 *model.PayRefund) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveRefund", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveRefund indicates an expected call of SaveRefund.
func (mr *MockRefundRepositoryMockRecorder) SaveRefund(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveRefund", reflect.TypeOf((*MockRefundRepository)(nil).SaveRefund), arg0, arg1)
}