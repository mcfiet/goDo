package model

type Todo struct {
	ID          int    `json:"id"`
	Name        string `json:"blod_name"`
	Description string `json:"blog_description,omitempty"`
}
