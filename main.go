package main

import (
	"fmt"

	db "users/grpc/client/database"
	"users/grpc/client/routes"
)

func main() {
	db.ConectaBanco()
	fmt.Println("Iniciando Servidor... ")
	routes.HandleRequest()
}
