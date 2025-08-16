package repository


import (
	"log"
	"time"
	//"os"
	"gorm.io/gorm"
	"github.com/tupap1/gologin/server/models"
	"gorm.io/driver/postgres"

)

func InitDatabase() (*gorm.DB, error) {
	dsn := "host=aws-0-sa-east-1.pooler.supabase.com user=postgres.pwrlfwrdsfoslavbyvde password=@Ndres137 dbname=postgres port=6543 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
