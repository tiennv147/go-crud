package model

import (
	"fmt"
	"net/http"
)

func AllTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "show all record")
}

func NewTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create record")
}
func UpdateTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "update record")
}
func DeleteTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete record")
}
