package config

import (
	"dheepssupreme/go-microservis.git/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// InitDB menginisialisasi koneksi ke database
func InitDB() {
	// Ambil path database dari environment variable atau gunakan default
	dsn := os.Getenv("DB_PATH")
	if dsn == "" {
		dsn = "./data/userdb.db" // Lokasi default jika DB_PATH tidak diset
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected")

	// Lakukan migrasi database
	MigrateDB()
}

// MigrateDB melakukan migrasi tabel ke database
func MigrateDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")
}
