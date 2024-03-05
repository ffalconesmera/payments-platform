package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBCon struct {
	*gorm.DB
}

type IDBCon interface {
	GetDatabase(host, port, name, user, password string) *DBCon
}

func NewDatabaseConnection(host, port, name, user, password string) *DBCon {
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

	return &DBCon{DB: db}
}

func (d *DBCon) Close() {
	sqlDB, err := d.DB.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get sql db: %s", err))
	}
	sqlDB.Close()
}
