package models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Phone_number    string `json:phone_number`
	Name            string `json:name`
	Email           string `json:email`
	Profile_picture string `json:profile_picture`
}
