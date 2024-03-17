package router

import (
	"net/http"
	repository "pitch-hunt/src/adapters/gateways/repository"
	database "pitch-hunt/src/adapters/primary/database"
	user_handler "pitch-hunt/src/adapters/primary/handlers/user"
	middleware "pitch-hunt/src/adapters/primary/middleware"

	"github.com/gorilla/mux"
)

func SetupUserRouting(r *mux.Router) {
	userRepository := repository.NewUserRepository(database.GetDB())
	authenticationMiddleware := middleware.NewAuthenticationMiddleware(userRepository)

	r.Handle(
		"/api/me",
		authenticationMiddleware.Handle(http.HandlerFunc(user_handler.NewMeHandler().Handle)),
	).Methods("GET")
}
