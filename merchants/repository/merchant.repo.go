package repository

import (
	"context"

	"github.com/ffalconesmera/payments-platform/merchants/database"
	"github.com/ffalconesmera/payments-platform/merchants/model"
)

// MerchantRepository is an interface to retrieve and create merchants
type MerchantRepository interface {
	FindMerchantById(ctxt context.Context, id string) (*model.PayMerchant, bool, error)
	FindMerchantByCode(ctxt context.Context, merchantCode string) (*model.PayMerchant, bool, error)
	CreateMerchant(ctxt context.Context, merchant *model.PayMerchant) error
}

type merchantRepositoryImpl struct {
	db *database.DBCon
}

func NewMerchantRepository(db *database.DBCon) MerchantRepository {
	db.AutoMigrate(&model.PayMerchant{})

	return &merchantRepositoryImpl{db: db}
}

// FindMerchantById: retrieve a merchant data by code
func (c *merchantRepositoryImpl) FindMerchantById(ctxt context.Context, id string) (*model.PayMerchant, bool, error) {
	whereMerchant := model.PayMerchant{UUID: id}
	var merchant model.PayMerchant
	result := c.db.WithContext(ctxt).Omit("User").Where(whereMerchant).Find(&merchant)
	return &merchant, result.RowsAffected > 0, result.Error
}

// FindMerchantByCode: retrieve a merchant data by code
func (c *merchantRepositoryImpl) FindMerchantByCode(ctxt context.Context, merchantCode string) (*model.PayMerchant, bool, error) {
	whereMerchant := model.PayMerchant{MerchantCode: merchantCode}
	var merchant model.PayMerchant
	result := c.db.WithContext(ctxt).Omit("User").Where(whereMerchant).Find(&merchant)
	return &merchant, result.RowsAffected > 0, result.Error
}

// CreateMerchant: store a new merchant
func (c *merchantRepositoryImpl) CreateMerchant(ctxt context.Context, merchant *model.PayMerchant) error {
	err := c.db.WithContext(ctxt).Create(&merchant)
	return err.Error
}
