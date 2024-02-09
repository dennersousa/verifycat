package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"verifycat/validate"
)

func main() {
	router := gin.Default()

	// Rota para manipulação de validações (aceita GET e POST)
	router.Any("/validate", validate.ValidateHandler)

	// Iniciar o servidor na porta 8080
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
