package route2

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIHandler interface {
	RegisterRoute(*mux.Router)
	NewErrResponse(http.ResponseWriter, string)
	QueryParse (http.ResponseWriter, *http.Request) *Query
	Posts (http.ResponseWriter, *http.Request)
}

type apiHandler struct {
	service Service
}

func NewAPIHandler(s Service) APIHandler {
	return &apiHandler{
		s,
	}
}
func (handler *apiHandler) RegisterRoute(r *mux.Router) {
	r.HandleFunc("/api/posts", handler.Posts).Methods("GET")
}

func (handler *apiHandler) NewErrResponse(w http.ResponseWriter, msg string) {
	newResponse := errResponse{
		Error: msg,
	}
	jsonNewResp, err := json.Marshal(newResponse)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.WriteHeader(400)
	w.Write(jsonNewResp)
}

func (handler *apiHandler) QueryParse (w http.ResponseWriter, r *http.Request) *Query{
	query := Query {
		r.URL.Query().Get("tags"),
		r.URL.Query().Get("sortBy"),
		r.URL.Query().Get("direction"),
	}
	// check if tags is empty,
	// return error response if empty
	if query.Tags == "" {
		handler.NewErrResponse(w, "Tags parameter is required")
		return nil
	}
	// check if sortBy is empty, then set to default id
	if query.SortBy == "" {
		query.SortBy = "id"
	} else {
		// check if sortBy is "id", "reads", "likes", or "popularity",
		// return error response if not
		sliceS := []string{"id","reads","likes","popularity"}
		if handler.service.Find(sliceS, query.SortBy) == false {
			handler.NewErrResponse(w, "SortBy parameter is invalid")
			return nil
		}
	}
	// check if direction is empty, then set to default asc
	if query.Direction == "" {
		query.Direction = "asc"
	} else {
		// check if direction is "asc", or "desc",
		// return error response if not
		sliceD := []string{"asc","desc"}
		if handler.service.Find(sliceD, query.Direction) == false {
			handler.NewErrResponse(w, "Direction parameter is invalid")
			return nil
		}
	}
	return &query
}

func (handler *apiHandler) Posts (w http.ResponseWriter, r *http.Request) {
	query := handler.QueryParse(w, r)
	if query == nil {
		return
	}
	// fetch data with tags and combine&de-duplicate if there are more than one tag
	data := handler.service.FetchData(query.Tags)
	// sort data
	result := handler.service.SortBy(query.SortBy, query.Direction, data)
	// write response
	jsonNewResp, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(200)
	w.Write(jsonNewResp)
}
