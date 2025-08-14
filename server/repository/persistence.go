package repository


import (
	"log"
	"os"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/tupap1/gologin/server/models"
)

func initDatabase() (*gorm.DB, error) {
	os.Remove("test.db")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	log.Println("Migración de todas las tablas completada.")
	return db, nil
}
