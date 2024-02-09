package main

import (
	"github.com/gatinhodev/verifycat/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Rota para manipulação de validações
	router.POST("/validate", validate.ValidateHandler)

	// Iniciar o servidor na porta 8080
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
