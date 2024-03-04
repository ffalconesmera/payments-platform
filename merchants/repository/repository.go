package repository

import (
	"github.com/ffalconesmera/payments-platform/merchants/database"
	"gorm.io/gorm"
)

// Repository is an interface por manage transactions
type Repository interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB)
	RollbackTransaction(tx *gorm.DB)
}

type RepositoryImpl struct {
	db database.DatabaseConnection
}

func (r *RepositoryImpl) BeginTransaction() *gorm.DB {
	return r.db.GetDatabase().Begin()
}

func (r *RepositoryImpl) CommitTransaction(tx *gorm.DB) {
	tx.Commit()
}

func (r *RepositoryImpl) RollbackTransaction(tx *gorm.DB) {
	tx.Rollback()
}
