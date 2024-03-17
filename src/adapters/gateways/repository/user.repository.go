package repository

import (
	"errors"
	"pitch-hunt/src/hexagon/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.First(&user, "email = ?", email)

	if err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return user, err.Error
	}

	return user, nil
}

func (r *UserRepository) Create(user models.User) (*models.User, error) {
	err := r.db.Create(&models.User{Id: user.Id, Username: user.Username, Email: user.Email, Password: user.Password})
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (r *UserRepository) FindById(id string) (*models.User, error) {
	user := &models.User{}
	err := r.db.First(&user, "id = ?", id)

	if err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return user, err.Error
	}

	return user, nil
}
