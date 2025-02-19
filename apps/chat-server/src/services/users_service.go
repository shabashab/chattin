package services

import (
	"github.com/shabashab/chattin/apps/chat-server/src/database/models"
	"gorm.io/gorm"
)

type UsersService struct {
	db *gorm.DB
}

func NewUsersService(db *gorm.DB) (*UsersService) {
	return &UsersService{
		db: db,
	}
}

func (s UsersService) FindUserById(id uint) (*models.User, error) {
	user := &models.User{}

	result := s.db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}