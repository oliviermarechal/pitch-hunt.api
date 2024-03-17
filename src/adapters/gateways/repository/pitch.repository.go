package repository

import (
	"pitch-hunt/src/hexagon/models"

	"gorm.io/gorm"
)

type PitchRepository struct {
	db *gorm.DB
}

func NewPitchRepository(db *gorm.DB) *PitchRepository {
	return &PitchRepository{db: db}
}

func (r *PitchRepository) Create(pitch models.Pitch) (*models.Pitch, error) {
	err := r.db.Create(&models.Pitch{Id: pitch.Id, UserId: pitch.UserId, Title: pitch.Title, Description: pitch.Description})
	if err.Error != nil {
		return nil, err.Error
	}

	return &pitch, nil
}

func (r *PitchRepository) GetMostRecentPitches(limit int) ([]models.Pitch, error) {
	var pitches []models.Pitch
	err := r.db.Order("created_at desc").Limit(limit).Find(&pitches)
	if err.Error != nil {
		return nil, err.Error
	}

	return pitches, nil
}
