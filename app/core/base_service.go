package core

import (
	"github.com/imufiid/goabsen/app/modules/user"
)

type BaseService struct {
	UserService user.Service
}

func InitService(repository *BaseRepository) *BaseService {
	return &BaseService{
		UserService: user.NewService(repository.UserRepo),
	}
}
