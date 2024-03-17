package hexagon_repository

import (
	"pitch-hunt/src/hexagon/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	Create(user models.User) (*models.User, error)
	FindById(id string) (*models.User, error)
}
