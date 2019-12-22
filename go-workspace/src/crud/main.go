package main

import (
	M "crud/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/get", M.AllTest).Methods("GET")
	myRouter.HandleFunc("/create", M.NewTest).Methods("POST")
	myRouter.HandleFunc("/update", M.UpdateTest).Methods("PUT")
	myRouter.HandleFunc("/delete", M.DeleteTest).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Test")
	handleRequests()
}
