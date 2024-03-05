package repository

import (
	"github.com/ffalconesmera/payments-platform/payments/database"
	"github.com/ffalconesmera/payments-platform/payments/model"
)

// CustomerRepository is an interface to retrieve and create customers
type CustomerRepository interface {
	FindCustomerById(id string) (*model.PayCustomer, bool, error)
	CreateCustomer(customer *model.PayCustomer) error
	Repository
}

type customerRepositoryImpl struct {
	db database.DatabaseConnection
	RepositoryImpl
}

func NewCustomerRepository(db database.DatabaseConnection) *customerRepositoryImpl {
	db.GetDatabase().AutoMigrate(&model.PayCustomer{})

	return &customerRepositoryImpl{db: db}
}

// FindCustomerById: retrieve a customer data by code
func (m *customerRepositoryImpl) FindCustomerById(id string) (*model.PayCustomer, bool, error) {
	whereCustomer := model.PayCustomer{UUID: id}
	var customer model.PayCustomer
	result := m.db.GetDatabase().Where(whereCustomer).Find(&customer)
	return &customer, result.RowsAffected > 0, result.Error
}

// CreateCustomer: store a new customer
func (m *customerRepositoryImpl) CreateCustomer(customer *model.PayCustomer) error {
	err := m.db.GetDatabase().Create(&customer)
	return err.Error
}
