package repository

import (
	"github.com/ffalconesmera/payments-platform/merchants/database"
	"github.com/ffalconesmera/payments-platform/merchants/model"
)

// MerchantRepository is an interface to retrieve and create merchants
type MerchantRepository interface {
	FindMerchantById(id string) (*model.PayMerchant, bool, error)
	FindMerchantByCode(merchantCode string) (*model.PayMerchant, bool, error)
	CreateMerchant(merchant *model.PayMerchant) error
	Repository
}

type merchantRepositoryImpl struct {
	db database.DatabaseConnection
	RepositoryImpl
}

func NewMerchantRepository(db database.DatabaseConnection) *merchantRepositoryImpl {
	db.GetDatabase().AutoMigrate(&model.PayMerchant{})

	return &merchantRepositoryImpl{db: db}
}

// FindMerchantById: retrieve a merchant data by code
func (m *merchantRepositoryImpl) FindMerchantById(id string) (*model.PayMerchant, bool, error) {
	whereMerchant := model.PayMerchant{UUID: id}
	var merchant model.PayMerchant
	result := m.db.GetDatabase().Omit("User").Where(whereMerchant).Find(&merchant)
	return &merchant, result.RowsAffected > 0, result.Error
}

// FindMerchantByCode: retrieve a merchant data by code
func (m *merchantRepositoryImpl) FindMerchantByCode(merchantCode string) (*model.PayMerchant, bool, error) {
	whereMerchant := model.PayMerchant{MerchantCode: merchantCode}
	var merchant model.PayMerchant
	result := m.db.GetDatabase().Omit("User").Where(whereMerchant).Find(&merchant)
	return &merchant, result.RowsAffected > 0, result.Error
}

// CreateMerchant: store a new merchant
func (m *merchantRepositoryImpl) CreateMerchant(user *model.PayMerchant) error {
	err := m.db.GetDatabase().Create(&user)
	return err.Error
}
