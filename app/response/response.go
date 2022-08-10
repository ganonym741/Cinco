package response

import "gitlab.com/cinco/app/model"

type (
	RegisterResponse struct {
		Messages string     `json:"message"`
		Data     model.User `json:"data"`
	}

	LoginResponse struct {
		Status   string `json:"status"`
		Messages string `json:"message"`
		Token    string `json:"token"`
	}

	LogoutResponse struct {
		Status   string `json:"status"`
		Messages string `json:"message"`
		Token    string `json:"token"`
	}
)
