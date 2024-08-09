package models

import "gorm.io/gorm"

type Todo struct {
	TaskNumber  int
	Description string
	Completed   string
	DayCount    int
	UserName    string
}

type User struct {
	gorm.Model
	UserName string
	Password string
}
