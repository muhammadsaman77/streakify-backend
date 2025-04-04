package config

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDB()*gorm.DB{
	dsn := "host=localhost user=postgres password=secret dbname=steakify-db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err:= gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	return db
}

func SetupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a mock database connection", err)
		return nil, nil, err
	}

	// Membuat koneksi GORM dari mock SQL
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a GORM database", err)
		return nil, nil, err
	}

	return gormDB, mock, nil
}