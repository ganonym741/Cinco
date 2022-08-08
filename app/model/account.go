package model

import (
	"gorm.io/gorm"
)

type Account struct {
	UserId    string `gorm:"type:uuid" json:"userid"`
	AccountId string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"AccountId"`
	Balance   int    `json:"balance"`
	gorm.Model
}
