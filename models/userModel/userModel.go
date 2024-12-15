package userModel

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
