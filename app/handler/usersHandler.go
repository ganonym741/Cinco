package handler

import (
	"gitlab.com/cinco/app/service/interfaces"
)

type CincoUser interface {
	UserRegister()
	UserLogin()
	UserLogout()
	UserProfile()
}

type UserHandler struct {
	UserService interfaces.UserServiceInterface
}

func (u UserHandler) UserRegister() {
	//TODO implement me
	panic("implement me")
}

func (u UserHandler) UserLogin() {
	//TODO implement me
	panic("implement me")
}

func (u UserHandler) UserLogout() {
	//TODO implement me
	panic("implement me")
}

func (u UserHandler) UserProfile() {
	//TODO implement me
	panic("implement me")
}

/*func (h Handler) UserProfile(c *fiber.Ctx) error {
	ctx := context.Background()
	params := c.Params("userId")
	data, err := h.service.GetUserDetail(ctx, params)
	if err != nil {
		return c.Status(201).
			JSON(fiber.Map{"status": "failed", "message": "Data not found", "data": nil})
	}

	return c.Status(201).
		JSON(fiber.Map{"status": "success", "message": "User data retrieved", "data": data})
}*/

func NewUserHandler(service interfaces.UserServiceInterface) CincoUser {
	return &UserHandler{
		UserService: service,
	}
}
