package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId      string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"userid"`
	Username    string    `gorm:"type:unique" json:"username"`
	Password    string    `json:"password"`
	Fullname    string    `json:"fullname"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateofbirth"`
	gorm.Model
}
