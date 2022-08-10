package model

import (
	"gorm.io/gorm"
)

type Account struct {
	UserId    string `gorm:"type:uuid" json:"user_id"`
	AccountId string `gorm:"type:uuid;primary_key" json:"AccountId"`
	Balance   int    `json:"balance"`
	gorm.Model
}
