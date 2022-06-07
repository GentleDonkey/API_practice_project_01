package server

import (
	"github.com/gorilla/mux"
)

func SetServer() (r *mux.Router) {
	// use gorilla mux to set server
	router := mux.NewRouter().StrictSlash(true)
	return router
}
