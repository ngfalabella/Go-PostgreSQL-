package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title 	string `gorm:"not null;unique_index"`
	Desc 	string
	Done 	bool	`gorm:"default:false"`
	UserId 	uint
}