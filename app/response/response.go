package response

import (
	"time"

	"gitlab.com/cinco/app/model"
)

type (
	RegisterResponse struct {
		Messages string     `json:"message"`
		Data     model.User `json:"data"`
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
