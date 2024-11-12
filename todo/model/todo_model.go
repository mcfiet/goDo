package model

import "github.com/mcfiet/goDo/model"

type Todo struct {
	model.Base
	Name        string `json:"name"`
	Description string `json:"description"`
}
