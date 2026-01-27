package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB:", err)
	}

	//WAJIB untuk Railway + Supabase
	sqlDB.SetMaxOpenConns(10)                  // maksimal koneksi aktif
	sqlDB.SetMaxIdleConns(5)                   // koneksi idle
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // refresh koneksi

	return db
}
