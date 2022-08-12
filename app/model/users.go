package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         string    `gorm:"type:uuid;primary_key" json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Fullname   string    `json:"fullname,omitempty"`
	Email      string    `json:"email"`
	BirthDate  time.Time `json:"birthdate,omitempty"`
	Domicile   string    `json:"domicile,omitempty"`
	Occupation string    `json:"occupation,omitempty"`
	Status     bool      `json:"status"`
	gorm.Model
}
