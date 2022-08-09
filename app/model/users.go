package model

import (
	"time"
)

type User struct {
	Id          string     `gorm:"type:uuid;primary_key" json:"id"`
	Username    string     `json:"username"`
	Fullname    string     `json:"fullname"`
	Password    string     `json:"password"`
	Email       string     `json:"email"`
	Status      bool       `json:"status"`
	DateOfBirth time.Time  `json:"dateofbirth"`
	Domicile    string     `json:"domicile"`
	Occupation  string     `json:"occupation"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   int        `json:"updated_by"`
}
