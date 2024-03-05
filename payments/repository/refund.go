package repository

import (
	"github.com/ffalconesmera/payments-platform/payments/database"
	"github.com/ffalconesmera/payments-platform/payments/model"
)

// MerchantRepository is an interface to retrieve and create refunds
type RefundRepository interface {
	FindRefundById(id string) (*model.PayRefund, bool, error)
	CreateRefund(refund *model.PayRefund) error
	SaveRefund(refund *model.PayRefund) error
	Repository
}

type refundRepositoryImpl struct {
	db database.DatabaseConnection
	RepositoryImpl
}

func NewRefundRepository(db database.DatabaseConnection) *refundRepositoryImpl {
	db.GetDatabase().AutoMigrate(&model.PayRefund{})
	return &refundRepositoryImpl{db: db}
}

// FindRefundById: retrieve refund data by id
func (u *refundRepositoryImpl) FindRefundById(id string) (*model.PayRefund, bool, error) {
	whereRefund := model.PayRefund{UUID: id}
	var refund model.PayRefund
	result := u.db.GetDatabase().Where(whereRefund).Find(&refund)
	return &refund, result.RowsAffected > 0, result.Error
}

// CreateRefund: store a new refund
func (u *refundRepositoryImpl) CreateRefund(refund *model.PayRefund) error {
	err := u.db.GetDatabase().Create(&refund)
	return err.Error
}

// CreateRefund: update a refund
func (u *refundRepositoryImpl) SaveRefund(refund *model.PayRefund) error {
	whereRefund := model.PayRefund{UUID: refund.UUID}
	err := u.db.GetDatabase().Where(whereRefund).Save(&refund)
	return err.Error
}
