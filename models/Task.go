package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title string
	Desc string
	Done bool
	UserId uint
}