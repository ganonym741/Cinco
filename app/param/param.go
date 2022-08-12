package param

import _ "github.com/go-playground/validator/v10"

type (
	User struct {
		Id         string `gorm:"type:uuid;primary_key" json:"id"`
		Username   string `json:"username" validate:"required,min=5,unique,alpha"`
		Fullname   string `json:"fullname" validate:"required"`
		Password   string `json:"password" validate:"required,min=5"`
		Email      string `json:"email" validate:"required,email"`
		BirthDate  string `json:"birth_date"`
		Domicile   string `json:"domicile"`
		Occupation string `json:"occupation"`
	}

	Login struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
)
