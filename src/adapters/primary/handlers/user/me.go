package user_handler

import (
	"encoding/json"
	"net/http"

	"pitch-hunt/src/hexagon/models"
)

type MeHandler struct {
}

func NewMeHandler() *MeHandler {
	return &MeHandler{}
}

func (c *MeHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(*models.User)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}
