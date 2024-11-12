package model

import "github.com/mcfiet/goDo/utils"

type User struct {
	utils.Base
	Username string `json:"username"`
	Password string `json:"password"`
}
