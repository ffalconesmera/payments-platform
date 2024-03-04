package database

import (
	"fmt"
	"log"

	"github.com/ffalconesmera/payments-platform/merchants/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database: is an interface manage connections with database
type Database interface {
	GetDatabase() *gorm.DB
	InitDatabase(host, port, dbName, user, password string)
	Close()
}

type databaseConnection struct {
	db     *gorm.DB
	config config.Config
}

func NewDatabaseConnection(config config.Config) *databaseConnection {
	return &databaseConnection{config: config}
}

// GetDatabase: return instance of gorm.DB
func (db *databaseConnection) GetDatabase() *gorm.DB {
	return db.db
}

// InitDatabase: initialize database connection
func (d *databaseConnection) InitDatabase() {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		d.config.GetDatabaseHost(),
		d.config.GetDatabasePort(),
		d.config.GetDatabaseName(),
		d.config.GetDatabaseUser(),
		d.config.GetDatabasePassword(),
	)

	log.Println("connecting with database..")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(fmt.Sprintf("failed to connect database: %s", err.Error()), true)
	}

	log.Println("database connected..!")

	d.db = db
}

// Close: close database connection
func (d *databaseConnection) Close() {
	log.Println("closing database connection..")

	sqlDB, err := d.db.DB()
	if err != nil {
		log.Println(fmt.Sprintf("failed to get sql db: %s", err))
	}

	sqlDB.Close()
}
