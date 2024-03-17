package create_pitch

import (
	hexagon_repository "pitch-hunt/src/hexagon/gateways/repository"
	"pitch-hunt/src/hexagon/models"

	"github.com/google/uuid"
)

type CreatePitchUseCase struct {
	pitchRepository hexagon_repository.PitchRepository
}

func NewCreatePitchUseCase(pitchRepository hexagon_repository.PitchRepository) *CreatePitchUseCase {
	return &CreatePitchUseCase{
		pitchRepository: pitchRepository,
	}
}

func (c *CreatePitchUseCase) Handle(command CreatePitchCommand) (*models.Pitch, error) {
	pitch := models.CreatePitch(
		uuid.New().String(),
		command.UserId,
		command.VideoUrl,
		command.Title,
		command.Description,
	)

	saved, err := c.pitchRepository.Create(*pitch)
	if err != nil {
		return &models.Pitch{}, err
	}

	return saved, nil
}
