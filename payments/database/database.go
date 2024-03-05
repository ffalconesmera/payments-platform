package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var singleDatabase *gorm.DB

type DatabaseConnection interface {
	GetDatabase() *gorm.DB
	InitDatabase(host, port, name, user, password string)
	Close()
}

type databaseConnection struct{}

func NewDatabaseConnection() DatabaseConnection {
	return &databaseConnection{}
}

func (d *databaseConnection) GetDatabase() *gorm.DB {
	return singleDatabase
}

func (d *databaseConnection) InitDatabase(host, port, name, user, password string) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host,
		port,
		name,
		user,
		password,
	)

	log.Println("connecting with database..")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(fmt.Sprintf("failed to connect database: %s", err.Error()), true)
	}

	log.Println("database connected..!")

	singleDatabase = db
}

func (d *databaseConnection) Close() {
	sqlDB, err := d.GetDatabase().DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get sql db: %s", err))
	}
	sqlDB.Close()
}
