package model

import (
	"gorm.io/gorm"
)

type Account struct {
	Id      string `gorm:"type:uuid;primary_key" json:"Id,omitempty"`
	UserId  string `gorm:"type:uuid" json:"userid,omitempty"`
	Balance int    `json:"balance,omitempty"`
	gorm.Model
}
