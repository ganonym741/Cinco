package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	utilities "gitlab.com/cinco/utils"
	"time"

	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/param"
	"gitlab.com/cinco/app/repository"
	"gitlab.com/cinco/app/response"
	"gitlab.com/cinco/pkg/postgres"
)

type Service struct {
	repository repository.Repository
}

func NewService() Service {
	return Service{
		repository: repository.Repository{
			Db: postgres.ConnectDB(),
		},
	}
}

func (s Service) UserRegister(ctx *fiber.Ctx, params *param.User) (*response.RegisterResponse, error) {
	params.Id = uuid.New().String()
	params.Password, _ = utilities.GeneratePassword(params.Password)
	date, _ := time.Parse(utilities.LayoutFormat, params.BirthDate)

	createdRegister := model.User{
		Id:         params.Id,
		Username:   params.Username,
		Fullname:   params.Fullname,
		Password:   params.Password,
		Email:      params.Email,
		BirthDate:  date,
		Domicile:   params.Domicile,
		Occupation: params.Occupation,
	}

	err := s.repository.UserRegister(ctx, createdRegister)
	if err != nil {
		return nil, err
	}

	return &response.RegisterResponse{
		Messages: "Register Success",
		Data:     createdRegister,
	}, nil
}

func (s Service) GetUserDetail(ctx *fiber.Ctx, userid string) (*model.User, error) {
	var data model.User
	err := s.repository.GetUserDetail(ctx, &data, userid)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s Service) UserLogin(ctx *fiber.Ctx, params *param.Login) (*response.LoginResponse, error) {
	result, err := s.repository.GetUserByIdentity(ctx, params.Identity)
	if err != nil {
		return nil, err
	}

	isMatch := utilities.ComparePasswords(result.Password, []byte(params.Password))
	if !isMatch {
		ctx.Status(403).JSON(fiber.Map{
			"status":  "failed",
			"message": "Wrong username & password",
			"data":    nil,
		})
	}

	token := utilities.CreateToken(result)

	return &response.LoginResponse{
		Status:   "success",
		Messages: "User data retrieved",
		Token:    token,
	}, nil

}
