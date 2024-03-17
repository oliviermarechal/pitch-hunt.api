package pitch_handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	hexagon_repository "pitch-hunt/src/hexagon/gateways/repository"
	list_pitch "pitch-hunt/src/hexagon/use-cases/pitch/query/list-pitch"
)

type ListPitchHandler struct {
	pitchRepository hexagon_repository.PitchRepository
}

func NewListPitchHander(pitchRepository hexagon_repository.PitchRepository) *ListPitchHandler {
	return &ListPitchHandler{
		pitchRepository: pitchRepository,
	}
}

func (c *ListPitchHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	query := list_pitch.NewListPitchQuery(c.pitchRepository)

	pitchs, err := query.Handle()

	if err != nil {
		resp := map[string]string{
			"message": err.Error(),
			"status":  strconv.Itoa(http.StatusBadRequest),
		}

		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(resp)
		return
	}

	json.NewEncoder(writer).Encode(pitchs)
}
