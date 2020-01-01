package main

import (
	"log"
	"net/http"
	s "service"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("api/test", s.getAll).Methods("GET")
	route.HandleFunc("api/test/{id}", s.getId).Methods("GET")
	route.HandleFunc("api/test/create", s.CreateTest).Methods("POST")
	route.HandleFunc("api/test/update", s.UpdateTest).Methods("PUT")
	route.HandleFunc("api/test/{id}", s.DeleteTest).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", route))
}
