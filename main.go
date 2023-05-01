package main

import (
	"fmt"

	"github.com/eduardotvn/projeto-api/envInitializer"
	"github.com/eduardotvn/projeto-api/src/postgres"
	"github.com/eduardotvn/projeto-api/src/server"
)

func init() {
	envInitializer.LoadEnvVar()
	postgres.DbConnect()
}

func main() {

	fmt.Println("Teste")

	server := server.StartServer()
	server.Run()
}
