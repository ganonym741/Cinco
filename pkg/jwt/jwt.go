package jwt

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	models "gitlab.com/cinco/app/model"
	"gitlab.com/cinco/configs"
)

func New() *JwtManager {
	return &JwtManager{}
}

type JwtManager struct{}

func (jwtManager *JwtManager) CreateToken(user *models.User) string {
	configs := configs.Config()

	claims := jwt.MapClaims{
		"userid": user.Id,
		"exp":    time.Now().Add(time.Duration(configs.Jwtconfig.Expired) * time.Second).Unix(),
	}

	unsignToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, _ := unsignToken.SignedString([]byte(configs.Jwtconfig.Secret))

	return token
}

func (jwtManager *JwtManager) GetUserId(c *fiber.Ctx) (UserId uuid.UUID) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	UserId, _ = uuid.Parse(claims["userid"].(string))
	return
}
