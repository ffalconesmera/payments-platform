package service

import (
	"context"
	"errors"
	externals_dto "github.com/ffalconesmera/payments-platform/payments/externals/dto"
	mock_repository "github.com/ffalconesmera/payments-platform/payments/mocks/repository"
	"github.com/ffalconesmera/payments-platform/payments/model"
	"github.com/ffalconesmera/payments-platform/payments/model/dto"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSingUp(t *testing.T) {
	t.Run("failed in method FindMerchantByCode", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		merchantRepository.EXPECT().FindMerchantByCode("").Return(nil, errors.New("merchant not found"))
		paymentService := NewPaymentService(nil, nil, nil, nil, merchantRepository)
		response, err := paymentService.CheckoutPayment(context.TODO(), "", dto.Payment{})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method CreateCustomer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchant := externals_dto.Merchant{}

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		refundRepository := mock_repository.NewMockRefundRepository(ctrl)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		merchantRepository.EXPECT().FindMerchantByCode("").Return(&merchant, nil)

		///println(err1)
		customerRepository := mock_repository.NewMockCustomerRepository(ctrl)
		customerRepository.EXPECT().CreateCustomer(context.TODO(), gomock.Any()).Return(errors.New("error saving customer"))
		paymentService := NewPaymentService(paymentRepository, refundRepository, customerRepository, bankRepository, merchantRepository)
		response, err := paymentService.CheckoutPayment(context.TODO(), "", dto.Payment{})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method CreatePayment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchant := externals_dto.Merchant{}

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		refundRepository := mock_repository.NewMockRefundRepository(ctrl)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		merchantRepository.EXPECT().FindMerchantByCode("").Return(&merchant, nil)

		///println(err1)
		customerRepository := mock_repository.NewMockCustomerRepository(ctrl)
		customerRepository.EXPECT().CreateCustomer(context.TODO(), gomock.Any()).Return(nil)
		paymentRepository.EXPECT().CreatePayment(context.TODO(), gomock.Any()).Return(errors.New("error saving payment"))
		paymentService := NewPaymentService(paymentRepository, refundRepository, customerRepository, bankRepository, merchantRepository)
		response, err := paymentService.CheckoutPayment(context.TODO(), "", dto.Payment{})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("successful", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchant := externals_dto.Merchant{}

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		refundRepository := mock_repository.NewMockRefundRepository(ctrl)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		merchantRepository.EXPECT().FindMerchantByCode("").Return(&merchant, nil)

		///println(err1)
		customerRepository := mock_repository.NewMockCustomerRepository(ctrl)
		customerRepository.EXPECT().CreateCustomer(context.TODO(), gomock.Any()).Return(nil)
		paymentRepository.EXPECT().CreatePayment(context.TODO(), gomock.Any()).Return(nil)
		paymentService := NewPaymentService(paymentRepository, refundRepository, customerRepository, bankRepository, merchantRepository)
		response, err := paymentService.CheckoutPayment(context.TODO(), "", dto.Payment{})

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}

func TestCheckPayment(t *testing.T) {
	t.Run("failed in method FindPaymentByCode", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), "").Return(nil, false, errors.New("payment not found"))
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.CheckPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method FindPaymentByCode is false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), "").Return(nil, false, nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.CheckPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed by merchantCode is different", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), "").Return(&model.PayTransaction{MerchantCode: "123"}, true, nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.CheckPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method FindCustomerById", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		customerRepository := mock_repository.NewMockCustomerRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), "").Return(&model.PayTransaction{MerchantCode: ""}, true, nil)
		customerRepository.EXPECT().FindCustomerById(context.TODO(), "").Return(nil, false, errors.New("error finding customer"))
		paymentService := NewPaymentService(paymentRepository, nil, customerRepository, nil, nil)
		response, err := paymentService.CheckPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method FindCustomerById is false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		customerRepository := mock_repository.NewMockCustomerRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), "").Return(&model.PayTransaction{MerchantCode: ""}, true, nil)
		customerRepository.EXPECT().FindCustomerById(context.TODO(), "").Return(nil, false, nil)
		paymentService := NewPaymentService(paymentRepository, nil, customerRepository, nil, nil)
		response, err := paymentService.CheckPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("successful", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		customerRepository := mock_repository.NewMockCustomerRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), "").Return(&model.PayTransaction{MerchantCode: ""}, true, nil)
		customerRepository.EXPECT().FindCustomerById(context.TODO(), "").Return(&model.PayCustomer{}, true, nil)
		paymentService := NewPaymentService(paymentRepository, nil, customerRepository, nil, nil)
		response, err := paymentService.CheckPayment(context.TODO(), "", "")

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}

func TestRefundPayment(t *testing.T) {
	t.Run("failed in method PaymentByCode", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(nil, false, errors.New("payment not found"))
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.RefundPayment(context.TODO(), "", "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method findPaymentByCode is false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(nil, false, nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.RefundPayment(context.TODO(), "", "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed by status payment is not pending or success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusFailure,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.RefundPayment(context.TODO(), "", "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed by merchantCode is different", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusFailure,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "123",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.RefundPayment(context.TODO(), "", "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method ProcessRefund", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusPending,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		bankRepository.EXPECT().ProcessRefund("", "").Return(nil, errors.New("bank return error"))
		paymentService := NewPaymentService(paymentRepository, nil, nil, bankRepository, nil)
		response, err := paymentService.RefundPayment(context.TODO(), "", "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method CreateRefund", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusPending,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		bankRefund := externals_dto.BankRefund{
			Status:    "",
			Code:      1000,
			Message:   "",
			Reference: "",
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		bankRepository.EXPECT().ProcessRefund("", "").Return(&bankRefund, nil)
		refundRepository := mock_repository.NewMockRefundRepository(ctrl)
		refundRepository.EXPECT().CreateRefund(context.TODO(), gomock.Any()).Return(errors.New("error saving refund"))
		paymentService := NewPaymentService(paymentRepository, refundRepository, nil, bankRepository, nil)
		response, err := paymentService.RefundPayment(context.TODO(), "", "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method SavePayment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusPending,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		bankRefund := externals_dto.BankRefund{
			Status:    "",
			Code:      1000,
			Message:   "",
			Reference: "",
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		bankRepository.EXPECT().ProcessRefund("", "").Return(&bankRefund, nil)
		refundRepository := mock_repository.NewMockRefundRepository(ctrl)
		refundRepository.EXPECT().CreateRefund(context.TODO(), gomock.Any()).Return(nil)
		paymentRepository.EXPECT().SavePayment(context.TODO(), gomock.Any()).Return(errors.New("error saving payment"))
		paymentService := NewPaymentService(paymentRepository, refundRepository, nil, bankRepository, nil)
		response, err := paymentService.RefundPayment(context.TODO(), "", "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed by response of bank code incorrect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusPending,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		bankRefund := externals_dto.BankRefund{
			Status:    "",
			Code:      2000,
			Message:   "",
			Reference: "",
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		bankRepository.EXPECT().ProcessRefund("", "").Return(&bankRefund, nil)
		paymentRepository.EXPECT().SavePayment(context.TODO(), gomock.Any()).Return(nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, bankRepository, nil)
		response, err := paymentService.RefundPayment(context.TODO(), "", "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("successul", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusPending,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		bankRefund2 := externals_dto.BankRefund{
			Status:    "",
			Code:      1000,
			Message:   "",
			Reference: "",
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		bankRepository.EXPECT().ProcessRefund("", "").Return(&bankRefund2, nil)
		refundRepository := mock_repository.NewMockRefundRepository(ctrl)
		refundRepository.EXPECT().CreateRefund(context.TODO(), gomock.Any()).Return(nil)
		paymentRepository.EXPECT().SavePayment(context.TODO(), gomock.Any()).Return(nil)
		paymentService := NewPaymentService(paymentRepository, refundRepository, nil, bankRepository, nil)
		response, err := paymentService.RefundPayment(context.TODO(), "", "", "")

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})

}

func TestProcessPayment(t *testing.T) {
	t.Run("failed in method FindPaymentByCode", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), "").Return(nil, false, errors.New("payment not found"))
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.ProcessPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method FindPaymentByCode is false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), "").Return(nil, false, nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.ProcessPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed by status payment is not pending", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusFailure,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, nil, nil)
		response, err := paymentService.ProcessPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method ProcessPayment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusPending,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		bankRepository.EXPECT().ProcessPayment("", "").Return(nil, errors.New("bank return error"))
		paymentService := NewPaymentService(paymentRepository, nil, nil, bankRepository, nil)
		response, err := paymentService.ProcessPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed by response of bank code incorrect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusPending,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		bankPayment := externals_dto.BankPayment{
			Status:    "",
			Code:      2000,
			Message:   "",
			Reference: "",
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		bankRepository.EXPECT().ProcessPayment("", "").Return(&bankPayment, nil)
		paymentRepository.EXPECT().SavePayment(context.TODO(), gomock.Any()).Return(nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, bankRepository, nil)
		response, err := paymentService.ProcessPayment(context.TODO(), "", "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("successful", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		modelMock := model.PayTransaction{
			UUID:                     "",
			PaymentCode:              "",
			Amount:                   0,
			Description:              "",
			Currency:                 "",
			Status:                   model.TransactionStatusPending,
			ExpirationProcess:        nil,
			NaturalExpirationProcess: "",
			FailureReason:            nil,
			BankReference:            nil,
			BankName:                 nil,
			MerchantCode:             "",
			CustomerUUID:             "",
			RefundUUID:               nil,
			BaseModel:                model.BaseModel{},
		}

		bankPayment := externals_dto.BankPayment{
			Status:    "",
			Code:      1000,
			Message:   "",
			Reference: "",
		}

		paymentRepository := mock_repository.NewMockPaymentRepository(ctrl)
		paymentRepository.EXPECT().FindPaymentByCode(context.TODO(), gomock.Any()).Return(&modelMock, true, nil)
		bankRepository := mock_repository.NewMockBankRepository(ctrl)
		bankRepository.EXPECT().ProcessPayment("", "").Return(&bankPayment, nil)
		paymentRepository.EXPECT().SavePayment(context.TODO(), gomock.Any()).Return(nil)
		paymentService := NewPaymentService(paymentRepository, nil, nil, bankRepository, nil)
		response, err := paymentService.ProcessPayment(context.TODO(), "", "")

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}
