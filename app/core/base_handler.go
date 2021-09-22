package core

import (
	"github.com/imufiid/goabsen/app/handler"
)

type BaseHandler struct {
	UserHandler handler.UserHandler
}

func InitHandler(service *BaseService) *BaseHandler {
	return &BaseHandler{
		UserHandler: handler.NewUserHandler(service.UserService),
	}
}
