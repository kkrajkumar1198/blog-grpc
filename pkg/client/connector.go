package client

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

func grpcConnector() (*grpc.ClientConn, error) {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	conn, err := grpc.Dial(serverAddress)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return conn, nil
}
