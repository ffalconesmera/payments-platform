package repository

import (
	"context"

	"github.com/ffalconesmera/payments-platform/merchants/database"
	"github.com/ffalconesmera/payments-platform/merchants/model"
)

// MerchantRepository is an interface to retrieve and create users
type UserRepository interface {
	FindUserByUsername(ctxt context.Context, username string) (*model.PayUser, bool, error)
	CreateUser(ctxt context.Context, user *model.PayUser) error
}

type userRepositoryImpl struct {
	db *database.DBCon
}

func NewUserRepository(db *database.DBCon) UserRepository {
	db.DB.AutoMigrate(&model.PayUser{})
	return &userRepositoryImpl{db: db}
}

// FindUserByUsername: retrieve a user data by code
func (c *userRepositoryImpl) FindUserByUsername(ctxt context.Context, username string) (*model.PayUser, bool, error) {
	whereUser := model.PayUser{Username: username}
	var user model.PayUser
	result := c.db.WithContext(ctxt).Where(whereUser).Find(&user)
	return &user, result.RowsAffected > 0, result.Error
}

// CreateUser: store a new user
func (c *userRepositoryImpl) CreateUser(ctxt context.Context, user *model.PayUser) error {
	err := c.db.WithContext(ctxt).Create(&user)
	return err.Error
}
