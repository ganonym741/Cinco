package model

import (
	"time"
)

type Cashflow struct {
	Id             string     `gorm:"type:uuid;primary_key" json:"id"`
	AccountId      string     `gorm:"type:uuid" json:"accountid,omitempty"`
	Type           string     `json:"type"`
	Description    string     `json:"description"`
	Amount         int        `json:"amount"`
	BalanceHistory int        `json:"balance_history"`
	IssuedAt       time.Time  `json:"issued_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
