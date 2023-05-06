package main

import (
	"fmt"
	"os"

	"github.com/eduardotvn/projeto-api/envInitializer"
	"github.com/eduardotvn/projeto-api/src/postgres"
	"github.com/eduardotvn/projeto-api/src/server"
)

func init() {
	envInitializer.LoadEnvVar()
	postgres.DBConnect()
	postgres.CreateTables()
}

func main() {

	port := os.Getenv("PORT")
	fmt.Println("Starting server in port:", port)
	server := server.StartServer()
	fmt.Println("Running server!")
	server.Run()
}
