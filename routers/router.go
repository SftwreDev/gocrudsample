package routers

import (
	"github.com/gorilla/mux"
	"gocrudsample/api"
	"gocrudsample/apps"
	"log"
	"net/http"
)

func InitializeRouter() {
	r := mux.NewRouter()

	// Page routers
	r.HandleFunc("/", apps.Homepage)

	// API endpoints
	r.HandleFunc("/actors", api.ActorGetListApi).Methods("GET")
	r.HandleFunc("/actors", api.ActorCreateAPI).Methods("POST")
	r.HandleFunc("/actors/details", api.ActorGetAPI).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/actors/update", api.ActorUpdateAPI).Methods("PUT").Queries("id", "{id}")
	r.HandleFunc("/actors/delete", api.ActorDeleteAPI).Methods("DELETE").Queries("id", "{id}")

	log.Fatal(http.ListenAndServe(":8080", r))
}
