package model

type Todo struct {
	ID          int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
