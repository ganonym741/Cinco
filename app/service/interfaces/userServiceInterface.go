package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/param"
	"gitlab.com/cinco/app/response"
)

type UserServiceInterface interface {
	FindByID(userUUID string) model.User
	Update(user model.User) error
	UserRegister(ctx *fiber.Ctx, params *param.User) (*model.User, error)
	UserLogout(ctx *fiber.Ctx, params string) (*response.LogoutResponse, error)
	UserLogin(ctx *fiber.Ctx, params *param.Login) (*response.LoginResponse, error)
	GetUserDetail(ctx *fiber.Ctx, userid string) (*model.User, error)
}
