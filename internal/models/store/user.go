package store

import (
	"ABCD/internal/models"
	"ABCD/internal/services"

	"gorm.io/gorm"
)

type UserStore struct {
	db           *gorm.DB
	passwordhash services.PasswordHash
}

func NewUserStore(db *gorm.DB, pwh services.PasswordHash) *UserStore {
	return &UserStore{
		db:           db,
		passwordhash: pwh,
	}
}

func (s *UserStore) CreateUser(email string, password string) error {

	hashedPassword, err := s.passwordhash.GenerateFromPassword(password)
	if err != nil {
		return err
	}

	return s.db.Create(&models.User{
		Email:    email,
		Password: hashedPassword,
	}).Error
}

func (s *UserStore) GetUser(email string) (*models.User, error) {

	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, err
}
