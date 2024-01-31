package client

import (
	"context"

	pb "github.com/kkrajkumar1198/blog-grpc/internal/blog/protos/bin"
)

func CreatePost(post *pb.Post) (string, error) {
	connection, err := grpcConnector()

	if err != nil {
		return "", err
	}
	serviceClient := pb.NewBlogServiceClient(connection)

	serverResponse, err := serviceClient.Create(context.Background(), &pb.Post{
		PostId:          post.PostId,
		Title:           post.Title,
		Content:         post.Content,
		Author:          post.Author,
		PublicationDate: post.PublicationDate,
		Tags:            post.Tags,
	})

	if err != nil {
		return "", err
	}

	response := serverResponse.PostId
	serverResponse.Response = response

	return serverResponse.Response, nil
}

func ReadPost(request *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	connection, err := grpcConnector()

	if err != nil {
		return nil, err
	}

	serviceClient := pb.NewBlogServiceClient(connection)

	serverResponse, err := serviceClient.Get(context.Background(), request)

	if err != nil {
		return nil, err
	}

	return serverResponse, nil
}

func DeletePost(request *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	connection, err := grpcConnector()

	if err != nil {
		return nil, err
	}

	serviceClient := pb.NewBlogServiceClient(connection)

	serverResponse, err := serviceClient.Delete(context.Background(), request)

	if err != nil {
		return nil, err
	}

	return serverResponse, nil
}
