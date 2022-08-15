package response

import (
	"time"
)

type (
	User struct {
		Id         string     `gorm:"type:uuid;primary_key" json:"id,omitempty"`
		Username   string     `json:"username,omitempty"`
		Fullname   string     `json:"fullname,omitempty"`
		Email      string     `json:"email,omitempty"`
		BirthDate  time.Time  `json:"birthdate,omitempty"`
		Domicile   string     `json:"domicile,omitempty"`
		Occupation string     `json:"occupation,omitempty"`
		Status     bool       `json:"status,omitempty"`
		CreatedAt  time.Time  `json:"created_at,omitempty"`
		UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	}

	RegisterResponse struct {
		Messages string `json:"message"`
		Data     User   `json:"data"`
	}

	LoginResponse struct {
		Status   string `json:"status"`
		Messages string `json:"message"`
		Token    string `json:"token"`
	}

	LogoutResponse struct {
		Status   string `json:"status"`
		Messages string `json:"message"`
		Token    string `json:"token"`
	}

	UserResponse struct {
		Fullname   string    `json:"fullname"`
		Email      string    `json:"email"`
		BirthDate  time.Time `json:"birth_date"`
		Domicile   string    `json:"domicile"`
		Occupation string    `json:"occupation"`
	}

	ProfileDetail struct {
		UserResponse
		AccountId string `json:"account_id"`
		Balance   string `json:"balance"`
	}
)
