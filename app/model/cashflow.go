package model

import (
	"time"

	"gorm.io/gorm"
)

type Cashflow struct {
	UserId         string    `gorm:"type:uuid" json:"userid"`
	CashflowId     string    `gorm:"type:uuid;primary_key" json:"AccountId"`
	Type           string    `json:"type"`
	Description    string    `json:"description"`
	Amount         int       `json:"amount"`
	BalanceHistory int       `json:"balance_history"`
	IssuedAt       time.Time `json:"issued_at"`
	gorm.Model
}
