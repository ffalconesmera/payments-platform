package repository

import (
	"github.com/ffalconesmera/payments-platform/payments/database"
	"github.com/ffalconesmera/payments-platform/payments/model"
)

// PaymentRepository is an interface to retrieve and create payments
type PaymentRepository interface {
	FindPaymentByCode(paymentCode string) (*model.PayTransaction, bool, error)
	CreatePayment(payment *model.PayTransaction) error
	SavePayment(payment *model.PayTransaction) error
	Repository
}

type paymentRepositoryImpl struct {
	db database.DatabaseConnection
	RepositoryImpl
}

func NewPaymentRepository(db database.DatabaseConnection) PaymentRepository {
	db.GetDatabase().AutoMigrate(&model.PayTransaction{})
	return &paymentRepositoryImpl{db: db}
}

// FindPaymentByCode: retrieve payment data by code
func (u *paymentRepositoryImpl) FindPaymentByCode(paymentCode string) (*model.PayTransaction, bool, error) {
	wherePayment := model.PayTransaction{PaymentCode: paymentCode}
	var payment model.PayTransaction
	result := u.db.GetDatabase().Where(wherePayment).Find(&payment)
	return &payment, result.RowsAffected > 0, result.Error
}

// CreatePayment: store a new payment
func (u *paymentRepositoryImpl) CreatePayment(refund *model.PayTransaction) error {
	err := u.db.GetDatabase().Create(&refund)
	return err.Error
}

// SavePayment: update a payment
func (u *paymentRepositoryImpl) SavePayment(payment *model.PayTransaction) error {
	wherePayment := model.PayTransaction{UUID: payment.UUID}
	err := u.db.GetDatabase().Where(wherePayment).Save(&payment)
	return err.Error
}
