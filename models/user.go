package models

import "gorm.io/gorm"

type user struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
