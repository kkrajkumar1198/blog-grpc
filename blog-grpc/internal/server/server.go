package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kkrajkumar1198/blog-grpc/internal/blog"
	pb "github.com/kkrajkumar1198/blog-grpc/internal/blog/protos/bin"
	"google.golang.org/grpc"
)

func StartServer() {
	port := os.Getenv("PORT")

	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Println("TCP ERROR" + err.Error())
		panic(err)
	}

	serve := grpc.NewServer()
	fmt.Println("SERVER RUNNING on: ", port)

	pb.RegisterBlogServiceServer(serve, &blog.Service{})

	if err = serve.Serve(listener); err != nil {
		log.Println("Server Not Started " + err.Error())
	}
}
