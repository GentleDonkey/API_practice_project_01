package route1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type httpResponse struct {
	Success bool `json:"success"`
}

func RegisterRoute(r *mux.Router) {
	r.HandleFunc("/api/ping", ping).Methods("GET")
}

func ping(w http.ResponseWriter, r *http.Request) {
	newResponse := httpResponse{
		Success: true,
	}
	jsonNewResp, err := json.Marshal(newResponse)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write(jsonNewResp)
	w.WriteHeader(200)
}
