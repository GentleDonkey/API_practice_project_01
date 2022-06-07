package main

import (
	"log"
	"net/http"
	"walnut/internal/route1"
	"walnut/internal/route2"
	"walnut/internal/server"
)

func main() {
	r := server.SetServer()
	// register a new router for Route1
	route1.RegisterRoute(r)
	// register a new router for Route2
	c := route2.NewAPIClient()
	s := route2.NewAPIService(c)
	h := route2.NewAPIHandler(s)
	h.RegisterRoute(r)
	// set TCP network address
	log.Fatal(http.ListenAndServe(":8000", r))
}



