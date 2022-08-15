package model

import (
	"time"
)

type Account struct {
	Id        string     `gorm:"type:uuid;primary_key" json:"Id,omitempty"`
	UserId    string     `gorm:"type:uuid" json:"userid,omitempty"`
	Balance   int        `json:"balance,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}
