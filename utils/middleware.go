package utilities

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
)

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func TokenVerify() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.GetRespHeader("Authorization")
		parts := strings.Split(token, " ")
		if token == "" {
			c.Status(401).JSON(fiber.Map{
				"status":  "failed",
				"message": "Unauthorized",
				"data":    nil,
			})
			return erro
		}
		claims, err := ExtractClaims(conf.Secret, parts[1])
		if err != nil {
			err := errs.ErrUnauthorized
			c.JSON(http.StatusUnauthorized, ErrorResponse(err))
			c.Abort()
			return
		}
		data := claims["Data"]
		result := map[string]interface{}{}
		encoded, _ := json.Marshal(data)
		json.Unmarshal(encoded, &result)
		for key, val := range result {
			c.Set(key, val)
		}
		c.Next()
	}
}
