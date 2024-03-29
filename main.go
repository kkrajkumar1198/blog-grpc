package main

import (
	"os"

	"github.com/joho/godotenv"
	database "github.com/kkrajkumar1198/blog-grpc/internal/databases"
	"github.com/kkrajkumar1198/blog-grpc/internal/server"
	"github.com/kkrajkumar1198/blog-grpc/pkg/client/http"
)

func main() {
	godotenv.Load()

	port := os.Getenv("HTTP_PORT")
	go http.Start(port)

	database.InitMigrations()
	server.StartServer()
}
