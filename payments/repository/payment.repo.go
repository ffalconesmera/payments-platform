package repository

import (
	"context"

	"github.com/ffalconesmera/payments-platform/payments/database"
	"github.com/ffalconesmera/payments-platform/payments/model"
)

// PaymentRepository is an interface to retrieve and create payments
type PaymentRepository interface {
	FindPaymentByCode(ctxt context.Context, paymentCode string) (*model.PayTransaction, bool, error)
	CreatePayment(ctxt context.Context, payment *model.PayTransaction) error
	SavePayment(ctxt context.Context, payment *model.PayTransaction) error
}

type paymentRepositoryImpl struct {
	db *database.DBCon
}

func NewPaymentRepository(db *database.DBCon) PaymentRepository {
	db.AutoMigrate(&model.PayTransaction{})
	return &paymentRepositoryImpl{db: db}
}

// FindPaymentByCode: retrieve payment data by code
func (c *paymentRepositoryImpl) FindPaymentByCode(ctxt context.Context, paymentCode string) (*model.PayTransaction, bool, error) {
	wherePayment := model.PayTransaction{PaymentCode: paymentCode}
	var payment model.PayTransaction
	result := c.db.WithContext(ctxt).Where(wherePayment).Find(&payment)
	return &payment, result.RowsAffected > 0, result.Error
}

// CreatePayment: store a new payment
func (c *paymentRepositoryImpl) CreatePayment(ctxt context.Context, refund *model.PayTransaction) error {
	err := c.db.WithContext(ctxt).Create(&refund)
	return err.Error
}

// SavePayment: update a payment
func (c *paymentRepositoryImpl) SavePayment(ctxt context.Context, payment *model.PayTransaction) error {
	wherePayment := model.PayTransaction{UUID: payment.UUID}
	err := c.db.WithContext(ctxt).Where(wherePayment).Save(&payment)
	return err.Error
}
