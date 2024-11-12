package model

import "github.com/mcfiet/goDo/model"

type User struct {
	model.Base
	Username string `json:"username"`
	Password string `json:"password"`
}
