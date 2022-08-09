package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         string    `gorm:"type:uuid;primary_key" json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Fullname   string    `json:"fullname"`
	Email      string    `json:"email"`
	BirthDate  time.Time `json:"birthdate"`
	Domicile   string    `json:"domicile"`
	Occupation string    `json:"occupation"`
	gorm.Model
}
