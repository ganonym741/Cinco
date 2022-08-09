package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/param"
	db "gitlab.com/cinco/pkg/postgres"
	utilities "gitlab.com/cinco/utils"
)

type CincoUser interface {
	UserRegister()
	UserLogin()
	UserLogout()
	UserProfile()
}

func UserRegister(c *fiber.Ctx) error {
	db := db.DB

	inputUser := new(param.User)
	inputUser.Id = uuid.New().String()
	err := c.BodyParser(inputUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}
	inputUser.Password, _ = utilities.GeneratePassword(inputUser.Password)
	fmt.Println(inputUser.Password)

	db.Create(&inputUser)

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "User data retrieved",
		"data":    inputUser,
	})
}

func UserLogin(c *fiber.Ctx) error {
	db := db.DB

	paramsLogin := new(param.Login)

	if err := c.BodyParser(paramsLogin); err != nil {
		return err
	}

	result := new(model.User)

	db.Where("username = ? or email = ?", paramsLogin.Identity, paramsLogin.Identity).Find(&result)

	isMatch := utilities.ComparePasswords(result.Password, []byte(paramsLogin.Password))
	if !isMatch {
		c.Status(403).JSON(fiber.Map{
			"status":  "failed",
			"message": "Wrong username & password",
			"data":    nil,
		})
	}

	token := utilities.CreateToken(result)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User data retrieved",
		"data":    token,
	})

}

//func UserLogout(c *fiber.Ctx) error {
//
//}
func UserProfile(c *fiber.Ctx) error {
	db := db.DB
	var user model.User

	userID := c.Query("id")

	db.Where("id = ?", userID).First(&user)

	if user.Id == "" {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "User data not found", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User data retrieved", "data": user})
}
