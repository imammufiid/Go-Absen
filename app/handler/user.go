package handler

import (
	"github.com/imufiid/goabsen/app/modules/user"
)

type UserHandler interface {
	// FunctionName(param dataType)
}

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

// Write code here, implementation interface