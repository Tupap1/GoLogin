package repository


import (
	"log"
	"time"
	//"os"
	"gorm.io/gorm"
	"github.com/tupap1/gologin/server/models"
	"gorm.io/driver/mysql"

)
dbpass := os.Getenv("STRING_DATABASE")
func InitDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("STRING_DATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sqlDB from gorm: %v", err)
	}
	

	sqlDB.SetMaxIdleConns(10) // Conexiones inactivas
	sqlDB.SetMaxOpenConns(100) // Máximo de conexiones abiertas
	sqlDB.SetConnMaxLifetime(time.Hour) // Tiempo de vida de la conexión

	log.Println("Successfully connected to the database!")


	err = db.AutoMigrate(&models.Products{})
	if err != nil {
		return nil, err
	}
	log.Println("Migración de todas las tablas completada.")
	return db, nil
}
