package core

import (
	"github.com/imufiid/goabsen/app/modules/user"
	"gorm.io/gorm"
)

type BaseRepository struct {
	UserRepo user.Repository
}

func InitRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		UserRepo: user.NewRepository(db),
	}
}
