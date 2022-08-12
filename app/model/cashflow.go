package model

import (
	"time"

	"gorm.io/gorm"
)

type Cashflow struct {
	Id             string    `gorm:"type:uuid;primary_key" json:"id"`
	AccountId      string    `gorm:"type:uuid" json:"accountid" db:"account_id,omitempty"`
	Type           string    `json:"type"`
	Description    string    `json:"description"`
	Amount         int       `json:"amount"`
	BalanceHistory int       `json:"balance_history"`
	IssuedAt       time.Time `json:"issued_at"`
	gorm.Model
}
