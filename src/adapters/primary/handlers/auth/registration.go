package auth_handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	hexagon_repository "pitch-hunt/src/hexagon/gateways/repository"
	registration "pitch-hunt/src/hexagon/use-cases/auth/command/registration"
)

type RegistrationHandler struct {
	userRepository hexagon_repository.UserRepository
}

func NewRegistrationHandler(userRepository hexagon_repository.UserRepository) *RegistrationHandler {
	return &RegistrationHandler{
		userRepository: userRepository,
	}
}

func (c *RegistrationHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	registrationDto := request.Context().Value("dto").(*registration.RegistrationDto)

	useCase := registration.NewRegistrationUseCase(c.userRepository)

	registrationResponse, err := useCase.Handle(
		registration.RegistrationCommand{
			Email:    registrationDto.Email,
			Password: registrationDto.Password,
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

	json.NewEncoder(writer).Encode(registrationResponse)
}
