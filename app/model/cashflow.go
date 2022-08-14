package model

import (
	"time"

	"gorm.io/gorm"
)

type Cashflow struct {
	Id             string    `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	AccountId      string    `gorm:"type:uuid" json:"accountid,omitempty"`
	Type           string    `json:"type,omitempty"`
	Description    string    `json:"description,omitempty"`
	Amount         int       `json:"amount,omitempty"`
	BalanceHistory int       `json:"balance_history,omitempty"`
	IssuedAt       time.Time `json:"issued_at,omitempty"`
	gorm.Model
}

type ResoponseCashflow struct {
	Description    string `json:"description,omitempty"`
	Amount         int    `json:"amount,omitempty"`
	BalanceHistory int    `json:"balance_history,omitempty"`
}
