package registration

import (
	"os"
	hexagon_repository "pitch-hunt/src/hexagon/gateways/repository"
	"pitch-hunt/src/hexagon/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationUseCase struct {
	UserRepository hexagon_repository.UserRepository
}

type RegistrationResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func NewRegistrationUseCase(UserRepository hexagon_repository.UserRepository) *RegistrationUseCase {
	return &RegistrationUseCase{UserRepository: UserRepository}
}

func (r *RegistrationUseCase) Handle(command RegistrationCommand) (RegistrationResponse, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(command.Password), 10)
	if err != nil {
		return RegistrationResponse{}, err
	}

	user := models.CreateUser(uuid.New().String(), command.Email, string(password))

	createdUser, err := r.UserRepository.Create((*user))
	if err != nil {
		return RegistrationResponse{}, err
	}

	claims := jwt.MapClaims{
		"id":  createdUser.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return RegistrationResponse{}, err
	}

	return RegistrationResponse{User: *createdUser, Token: signedToken}, nil
}
