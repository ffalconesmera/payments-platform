package repository

import (
	"context"

	"github.com/ffalconesmera/payments-platform/payments/database"
	"github.com/ffalconesmera/payments-platform/payments/model"
)

// MerchantRepository is an interface to retrieve and create refunds
type RefundRepository interface {
	FindRefundById(ctxt context.Context, id string) (*model.PayRefund, bool, error)
	CreateRefund(ctxt context.Context, refund *model.PayRefund) error
	SaveRefund(ctxt context.Context, refund *model.PayRefund) error
}

type refundRepositoryImpl struct {
	db *database.DBCon
}

func NewRefundRepository(db *database.DBCon) RefundRepository {
	db.AutoMigrate(&model.PayRefund{})
	return &refundRepositoryImpl{db: db}
}

// FindRefundById: retrieve refund data by id
func (c *refundRepositoryImpl) FindRefundById(ctxt context.Context, id string) (*model.PayRefund, bool, error) {
	whereRefund := model.PayRefund{UUID: id}
	var refund model.PayRefund
	result := c.db.WithContext(ctxt).Where(whereRefund).Find(&refund)
	return &refund, result.RowsAffected > 0, result.Error
}

// CreateRefund: store a new refund
func (c *refundRepositoryImpl) CreateRefund(ctxt context.Context, refund *model.PayRefund) error {
	err := c.db.WithContext(ctxt).Create(&refund)
	return err.Error
}

// CreateRefund: update a refund
func (c *refundRepositoryImpl) SaveRefund(ctxt context.Context, refund *model.PayRefund) error {
	whereRefund := model.PayRefund{UUID: refund.UUID}
	err := c.db.WithContext(ctxt).Where(whereRefund).Save(&refund)
	return err.Error
}
