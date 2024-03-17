package list_pitch

import (
	hexagon_repository "pitch-hunt/src/hexagon/gateways/repository"
	"pitch-hunt/src/hexagon/models"
)

type ListPitchQuery struct {
	pitchRepository hexagon_repository.PitchRepository
}

func NewListPitchQuery(pitchRepository hexagon_repository.PitchRepository) *ListPitchQuery {
	return &ListPitchQuery{
		pitchRepository: pitchRepository,
	}
}

func (h *ListPitchQuery) Handle() ([]models.Pitch, error) {
	pitches, err := h.pitchRepository.GetMostRecentPitches(10)
	if err != nil {
		return nil, err
	}

	return pitches, nil
}
