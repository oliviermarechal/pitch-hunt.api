package router

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	SetupAuthRouting(r)
	SetupPitchRouting(r)
	SetupUserRouting(r)
	return r
}
