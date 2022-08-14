package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         string    `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Password   string    `json:"password,omitempty"`
	Fullname   string    `json:"fullname,omitempty"`
	Email      string    `json:"email,omitempty"`
	BirthDate  time.Time `json:"birthdate,omitempty"`
	Domicile   string    `json:"domicile,omitempty"`
	Occupation string    `json:"occupation,omitempty"`
	Status     bool      `json:"status,omitempty"`
	gorm.Model
}
