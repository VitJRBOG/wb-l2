package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewConnection создает новое соединение с базой данных.
func NewConnection(dsn string) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("unable to connect to Postgres: %v\n", err)
	}

	return conn
}
