package db

import (
	"devoniq/goDo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (db *gorm.DB) {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Todo{})
	return db
}
