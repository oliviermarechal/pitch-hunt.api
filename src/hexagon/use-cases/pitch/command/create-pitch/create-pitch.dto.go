package create_pitch

type CreatePitchDto struct {
	Title       string `json:"title" validate:"required"`
	VideoUrl    string `json:"video_url" validate:"required`
	Description string `json:"description"`
}
