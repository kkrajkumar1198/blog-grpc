package client

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

func grpcConnector() (*grpc.ClientConn, error) {
	serverAddress := os.Getenv("BLOG_SERVER_ADDRESS") // Change to your specific environment variable name for the blog server address
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return conn, nil
}
