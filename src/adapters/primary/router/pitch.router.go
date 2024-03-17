package router

import (
	"net/http"
	repository "pitch-hunt/src/adapters/gateways/repository"
	database "pitch-hunt/src/adapters/primary/database"
	pitch_handler "pitch-hunt/src/adapters/primary/handlers/pitch"
	middleware "pitch-hunt/src/adapters/primary/middleware"
	create_pitch "pitch-hunt/src/hexagon/use-cases/pitch/command/create-pitch"

	"github.com/gorilla/mux"
)

func SetupPitchRouting(r *mux.Router) {
	validationMiddleware := middleware.NewValidationMiddleware()
	userRepository := repository.NewUserRepository(database.GetDB())
	pitchRepository := repository.NewPitchRepository(database.GetDB())

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(userRepository)

	r.Handle(
		"/api/pitch",
		authenticationMiddleware.Handle(
			validationMiddleware.Handle(create_pitch.CreatePitchDto{}, http.HandlerFunc(pitch_handler.NewCreatePitchHander(pitchRepository).Handle)),
		),
	).Methods("POST")
	r.HandleFunc(
		"/api/pitch",
		pitch_handler.NewListPitchHander(pitchRepository).Handle,
	).Methods("GET")
}
