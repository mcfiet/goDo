package model

import "github.com/mcfiet/goDo/utils"

type Todo struct {
	utils.Base
	Name        string `json:"name"`
	Description string `json:"description"`
}
