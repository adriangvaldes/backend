package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := SetupRoutes()

	fmt.Println("Servidor rodando na porta 4000...")
	if err := http.ListenAndServe(":4000", router); err != nil {
		panic(err)
	}
}
