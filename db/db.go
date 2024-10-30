package db

import (
	"devoniq/goDo/todo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

// Init initialisiert die Datenbank und speichert die Verbindung in der globalen `database`-Variable
func Init() {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migration des Models
	db.AutoMigrate(&model.Todo{})

	// Verbindung speichern
	database = db
}

// GetDB gibt die initialisierte DB-Verbindung zurück
func GetDB() *gorm.DB {
	return database
}

