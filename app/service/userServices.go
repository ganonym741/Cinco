package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/cinco/configs"
	utilities "gitlab.com/cinco/utils"
	"strings"
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

	err = utilities.SendMail(params.Email, "")

	return &response.RegisterResponse{
		Messages: "Register Success Check Your Email to Activated",
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
		return nil, errors.New("wrong email or password ")
	}

	if result.Status != true {
		return nil, errors.New("your account is deactive")
	}

	token := utilities.CreateToken(result)

	return &response.LoginResponse{
		Status:   "success",
		Messages: "User data retrieved",
		Token:    token,
	}, nil
}

func (s Service) UserLogout(ctx *fiber.Ctx, params string) (*response.LogoutResponse, error) {
	configs := configs.Config()
	token := strings.Split(ctx.Get("Authorization"), " ")
	claim, _ := utilities.ExtractClaims(configs.Jwtconfig.Secret, token[1])

	if claim["userid"] != params {
		var err error
		return nil, err
	}

	claim["exp"] = time.Now().Add(-time.Hour)
	return &response.LogoutResponse{
		Status:   "success",
		Messages: "logout",
		Token:    "",
	}, nil
}
