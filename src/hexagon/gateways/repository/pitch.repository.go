package hexagon_repository

import (
	"pitch-hunt/src/hexagon/models"
)

type PitchRepository interface {
	Create(user models.Pitch) (*models.Pitch, error)
	GetMostRecentPitches(limit int) ([]models.Pitch, error)
}
