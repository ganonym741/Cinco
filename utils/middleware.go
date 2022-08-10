package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/cinco/configs"
	"strings"
)

func ExtractClaims(secret, tokenStr string) (jwt.MapClaims, error) {
	hmacSecret := []byte(secret)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid JWT Token")
}

func TokenVerify() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		configs := configs.Config()
		token := ctx.GetReqHeaders()["Authorization"]
		parts := strings.Split(token, " ")
		if token == "" {
			return ctx.Status(401).JSON(fiber.Map{
				"status":  "failed",
				"message": "Unauthorized",
				"data":    nil,
			})
		}
		claims, err := ExtractClaims(configs.Jwtconfig.Secret, parts[1])
		if err != nil {
			return ctx.Status(401).JSON(fiber.Map{
				"status":  "failed",
				"message": "Unauthorized",
				"data":    nil,
			})
		}
		result := map[string]interface{}{}
		encoded, _ := json.Marshal(claims)
		json.Unmarshal(encoded, &result)
		for key, val := range result {
			valStr := fmt.Sprintf(`%v`, val)
			ctx.Set(key, valStr)
			//fmt.Println(key, valStr)
		}
		//fmt.Println("ini c get di atas =", c.GetReqHeaders())
		return ctx.Next()
	}
}

func Authorization(status bool) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		configs := configs.Config()
		token := strings.Split(ctx.Get("Authorization"), " ")
		claim, _ := ExtractClaims(configs.Jwtconfig.Secret, token[1])
		if claim["status"] != status {
			return ctx.Status(403).JSON(fiber.Map{
				"status":  "failed",
				"message": "Account not activate",
				"data":    nil,
			})
		}
		ctx.Next()
		return nil
	}
}
