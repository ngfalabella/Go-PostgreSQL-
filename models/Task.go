package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title 	string `gorm:"type:varchar(100);not null;unique_index"`
	Desc 	string
	Done 	bool	`gorm:"default:false"`
	UserId 	uint
}