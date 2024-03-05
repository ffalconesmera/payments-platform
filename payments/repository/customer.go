package repository

import (
	"context"

	"github.com/ffalconesmera/payments-platform/payments/database"
	"github.com/ffalconesmera/payments-platform/payments/model"
)

// CustomerRepository is an interface to retrieve and create customers
type CustomerRepository interface {
	FindCustomerById(ctxt context.Context, id string) (*model.PayCustomer, bool, error)
	CreateCustomer(ctxt context.Context, customer *model.PayCustomer) error
}

type customerRepositoryImpl struct {
	db *database.DBCon
}

func NewCustomerRepository(db *database.DBCon) CustomerRepository {
	db.AutoMigrate(&model.PayCustomer{})

	return &customerRepositoryImpl{db: db}
}

// FindCustomerById: retrieve a customer data by code
func (c *customerRepositoryImpl) FindCustomerById(ctxt context.Context, id string) (*model.PayCustomer, bool, error) {
	whereCustomer := model.PayCustomer{UUID: id}
	var customer model.PayCustomer
	result := c.db.WithContext(ctxt).Where(whereCustomer).Find(&customer)
	return &customer, result.RowsAffected > 0, result.Error
}

// CreateCustomer: store a new customer
func (c *customerRepositoryImpl) CreateCustomer(ctxt context.Context, customer *model.PayCustomer) error {
	err := c.db.WithContext(ctxt).Create(&customer)
	return err.Error
}
