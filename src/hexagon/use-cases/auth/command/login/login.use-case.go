package login

import (
	"errors"
	"os"
	hexagon_repository "pitch-hunt/src/hexagon/gateways/repository"
	"pitch-hunt/src/hexagon/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	UserRepository hexagon_repository.UserRepository
}

type LoginResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func NewLoginUseCase(UserRepository hexagon_repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{UserRepository: UserRepository}
}

func (l *LoginUseCase) Handle(command LoginCommand) (LoginResponse, error) {
	user, err := l.UserRepository.GetUserByEmail(command.Email)
	if err != nil {
		return LoginResponse{}, errors.New("authentication failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(command.Password))
	if err != nil {
		return LoginResponse{}, errors.New("authentication failed")
	}

	claims := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{User: *user, Token: signedToken}, nil
}
