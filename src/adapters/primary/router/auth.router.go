package router

import (
	"net/http"
	repository "pitch-hunt/src/adapters/gateways/repository"
	database "pitch-hunt/src/adapters/primary/database"
	auth_handler "pitch-hunt/src/adapters/primary/handlers/auth"
	middleware "pitch-hunt/src/adapters/primary/middleware"
	"pitch-hunt/src/hexagon/use-cases/auth/command/login"
	"pitch-hunt/src/hexagon/use-cases/auth/command/registration"

	"github.com/gorilla/mux"
)

func SetupAuthRouting(r *mux.Router) {
	validationMiddleware := middleware.NewValidationMiddleware()
	userRepository := repository.NewUserRepository(database.GetDB())

	r.Handle(
		"/api/login",
		validationMiddleware.Handle(login.LoginDto{}, http.HandlerFunc(auth_handler.NewLoginHandler(userRepository).Handle)),
	).Methods("POST")
	r.Handle(
		"/api/registration",
		validationMiddleware.Handle(registration.RegistrationDto{}, http.HandlerFunc(auth_handler.NewRegistrationHandler(userRepository).Handle)),
	).Methods("POST")
}
