package utilities

import (
	"github.com/golang-jwt/jwt/v4"
	models "gitlab.com/cinco/app/model"
	"gitlab.com/cinco/configs"
	"time"
)

func CreateToken(user *models.User) string {
	configs := configs.Config()

	claims := jwt.MapClaims{
		"userid": user.Id,
		"status": user.Status,
		"exp":    time.Now().Add(time.Duration(configs.Jwtconfig.Expired) * time.Second).Unix(),
	}

	unsignToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, _ := unsignToken.SignedString([]byte(configs.Jwtconfig.Secret))

	return token
}
