package repository

import (
	"github.com/ffalconesmera/payments-platform/merchants/database"
	"github.com/ffalconesmera/payments-platform/merchants/model"
)

// MerchantRepository is an interface to retrieve and create users
type UserRepository interface {
	FindUserByUsername(username string) (*model.PayUser, bool, error)
	CreateUser(user *model.PayUser) error
	Repository
}

type userRepositoryImpl struct {
	db database.DatabaseConnection
	RepositoryImpl
}

func NewUserRepository(db database.DatabaseConnection) *userRepositoryImpl {
	db.GetDatabase().AutoMigrate(&model.PayUser{})
	return &userRepositoryImpl{db: db}
}

// FindUserByUsername: retrieve a user data by code
func (u *userRepositoryImpl) FindUserByUsername(username string) (*model.PayUser, bool, error) {
	whereUser := model.PayUser{Username: username}
	var user model.PayUser
	result := u.db.GetDatabase().Where(whereUser).Find(&user)
	return &user, result.RowsAffected > 0, result.Error
}

// CreateUser: store a new user
func (u *userRepositoryImpl) CreateUser(user *model.PayUser) error {
	err := u.db.GetDatabase().Create(&user)
	return err.Error
}
