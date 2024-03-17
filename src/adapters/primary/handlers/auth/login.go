package auth_handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	hexagon_repository "pitch-hunt/src/hexagon/gateways/repository"
	login "pitch-hunt/src/hexagon/use-cases/auth/command/login"
)

type LoginHandler struct {
	userRepository hexagon_repository.UserRepository
}

func NewLoginHandler(userRepository hexagon_repository.UserRepository) *LoginHandler {
	return &LoginHandler{
		userRepository: userRepository,
	}
}

func (c *LoginHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	loginDto := request.Context().Value("dto").(*login.LoginDto)

	useCase := login.NewLoginUseCase(c.userRepository)

	createClientResponse, err := useCase.Handle(
		login.LoginCommand{
			Email:    loginDto.Email,
			Password: loginDto.Password,
		},
	)

	if err != nil {
		resp := map[string]string{
			"message": err.Error(),
			"status":  strconv.Itoa(http.StatusBadRequest),
		}

		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(resp)
		return
	}

	json.NewEncoder(writer).Encode(createClientResponse)
}
