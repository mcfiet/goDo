package model

type Todo struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
}
