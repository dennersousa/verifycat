package main

import (
	"github.com/gatinhodev/verifycat/validations"
	"net/http"
)

func main() {
	http.HandleFunc("/validate", validations.ValidateHandler)
	http.ListenAndServe(":8080", nil)
}
