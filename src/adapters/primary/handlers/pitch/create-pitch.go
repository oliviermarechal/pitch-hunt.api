package pitch_handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	hexagon_repository "pitch-hunt/src/hexagon/gateways/repository"
	"pitch-hunt/src/hexagon/models"
	create_pitch "pitch-hunt/src/hexagon/use-cases/pitch/command/create-pitch"
)

type CreatePitchHandler struct {
	pitchRepository hexagon_repository.PitchRepository
}

func NewCreatePitchHander(pitchRepository hexagon_repository.PitchRepository) *CreatePitchHandler {
	return &CreatePitchHandler{
		pitchRepository: pitchRepository,
	}
}

func (c *CreatePitchHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	createPitchDto := request.Context().Value("dto").(*create_pitch.CreatePitchDto)
	user := request.Context().Value("user").(*models.User)

	useCase := create_pitch.NewCreatePitchUseCase(c.pitchRepository)

	pitch, err := useCase.Handle(
		create_pitch.CreatePitchCommand{
			UserId:      user.Id,
			Title:       createPitchDto.Title,
			VideoUrl:    createPitchDto.VideoUrl,
			Description: createPitchDto.Description,
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

	json.NewEncoder(writer).Encode(pitch)
}
