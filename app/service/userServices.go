package service

import (
	"errors"

	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/cinco/configs"
	utilities "gitlab.com/cinco/utils"

	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/param"
	"gitlab.com/cinco/app/repository/interfaces"
	"gitlab.com/cinco/app/response"
	serviceInterface "gitlab.com/cinco/app/service/interfaces"
)

type UserService struct {
	userRepository interfaces.UserRepositoryInterface
}

func (u UserService) Update(user model.User) error {
	return u.userRepository.Update(user)
}

func (u UserService) FindByID(userUUID string) model.User {
	return u.userRepository.FindById(userUUID)
}

func (u UserService) UserRegister(ctx *fiber.Ctx, params *param.User) (*model.User, error) {
	params.Id = uuid.New().String()
	params.Password, _ = utilities.GeneratePassword(params.Password)
	activationLink := "Hallo," + params.Fullname + ", please actvate your account " +
		"<a href= \"http://" + configs.Config().Host + "/api/user/activation/" + params.Id + "\">here!</a>"

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

	err := u.userRepository.UserRegister(ctx, createdRegister)
	if err != nil {
		return nil, err
	}

	err = utilities.SendMail(params.Email, activationLink)
	if err != nil {
		return nil, err
	}

	return &createdRegister, nil
}

func (u UserService) GetUserDetail(ctx *fiber.Ctx, userid string) (*model.User, error) {
	var data model.User
	err := u.userRepository.GetUserDetail(ctx, &data, userid)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (u UserService) UserLogin(ctx *fiber.Ctx, params *param.Login) (*response.LoginResponse, error) {
	result, err := u.userRepository.GetUserByIdentity(ctx, params.Identity)
	if err != nil {
		return nil, err
	}

	isMatch := utilities.ComparePasswords(result.Password, []byte(params.Password))
	if !isMatch {
		return nil, errors.New("wrong email or password ")
	}

	if !result.Status {
		return nil, errors.New("your account is deactive")
	}

	token := utilities.CreateToken(result)

	return &response.LoginResponse{
		Status:   "success",
		Messages: "User data retrieved",
		Token:    token,
	}, nil
}

func (u UserService) UserLogout(ctx *fiber.Ctx, params string) (*response.LogoutResponse, error) {
	configs := configs.Config()
	token := strings.Split(ctx.Get("Authorization"), " ")
	claim, _ := utilities.ExtractClaims(configs.Jwtconfig.Secret, token[1])

	if claim["userid"] != params {
		var err error
		return nil, err
	}

	claim["exp"] = -1

	return &response.LogoutResponse{
		Status:   "success",
		Messages: "logout",
		Token:    "",
	}, nil
}

func NewUserService(repository interfaces.UserRepositoryInterface) serviceInterface.UserServiceInterface {
	return &UserService{
		userRepository: repository,
	}
}
