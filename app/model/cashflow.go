package model

import (
	"time"
)

type Cashflow struct {
	Id             string     `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	AccountId      string     `gorm:"type:uuid" json:"accountid,omitempty"`
	Type           string     `json:"type,omitempty"`
	Description    string     `json:"description,omitempty"`
	Amount         int        `json:"amount,omitempty"`
	BalanceHistory int        `json:"balance_history,omitempty"`
	IssuedAt       time.Time  `json:"issued_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

type ResoponseCashflow struct {
	Description    string `json:"description,omitempty"`
	Amount         int    `json:"amount,omitempty"`
	BalanceHistory int    `json:"balance_history,omitempty"`
}
