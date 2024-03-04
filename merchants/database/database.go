package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database is an interface manage connections with database
type Database interface {
	GetDatabase() *gorm.DB
	InitDatabase(host, port, dbName, user, password string)
	Close()
}

type databaseConnection struct {
	db *gorm.DB
}

func NewDatabaseConnection() *databaseConnection {
	return &databaseConnection{}
}

func (db *databaseConnection) GetDatabase() *gorm.DB {
	return db.db
}

func (d *databaseConnection) InitDatabase(host, port, dbName, user, password string) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host,
		port,
		dbName,
		user,
		password,
	)

	log.Println("connecting with database..")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(fmt.Sprintf("failed to connect database: %s", err.Error()), true)
	}

	log.Println("database connected..!")

	d.db = db
}

func (d *databaseConnection) Close() {
	log.Println("closing database connection..")

	sqlDB, err := d.db.DB()
	if err != nil {
		log.Println(fmt.Sprintf("failed to get sql db: %s", err))
	}

	sqlDB.Close()
}
