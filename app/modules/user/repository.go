package user

import "gorm.io/gorm"

type Repository interface {
	// FunctionName(param dataType)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

// Write code here, implementation interface