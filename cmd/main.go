package main

import (
	s "Go-Crud/go-crud/service"
	"fmt"
	"net/http"

	"tannv/go/src/github.com/gorilla/mux"
)

func main() {

	fmt.Println("Go run ...")

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/find", s.FindUser).Methods("GET")
	router.HandleFunc("/api/v1/user/getall", s.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/user/create", s.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update", s.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete", s.Delete).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
