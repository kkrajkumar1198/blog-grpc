package tests

import (
	"context"
	"testing"

	"github.com/kkrajkumar1198/blog-grpc/internal/blog"
	pb "github.com/kkrajkumar1198/blog-grpc/internal/blog/protos/bin"
)

func init() {
	blog.DataAccess = mockDataAccess{}
}

func TestCreateGRPC(t *testing.T) {
	pbPost := &pb.Post{
		PostId:          "1",
		Title:           "Sample Title",
		Content:         "Sample Content",
		Author:          "Sample Author",
		PublicationDate: "2024-01-31",
		Tags:            []string{"tag1", "tag2"},
	}

	service := blog.Service{}

	responseGrpc, err := service.CreatePost(context.Background(), pbPost)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	if responseGrpc.PostId == "" {
		t.Errorf("expected %v and got %v", pbPost.PostId, responseGrpc.PostId)
	}
}

func TestGetGRPC(t *testing.T) {
	request := &pb.GetPostRequest{
		PostId: "1",
	}

	rightPost := &pb.Post{
		PostId:          "1",
		Title:           "ed",
		Content:         "mart",
		Author:          "email",
		PublicationDate: "2024-01-31",
		Tags:            []string{"tag1", "tag2"},
	}

	wrongPost := &pb.Post{
		PostId:          "2",
		Title:           "mock2name",
		Content:         "mock2lastname",
		Author:          "no email",
		PublicationDate: "2024-01-31",
		Tags:            nil,
	}

	service := blog.Service{}

	result, err := service.GetPost(context.Background(), request)

	if err != nil {
		t.Errorf("Error %s", err)
	}

	if result.Post.PostId != rightPost.PostId {
		t.Errorf("expected %s and got %s", rightPost.PostId, result.Post.PostId)
	}

	if wrongPost.PostId == result.Post.PostId {
		t.Errorf("expected %s and got %s", wrongPost.PostId, result.Post.PostId)
	}
}

func TestDeleteGRPC(t *testing.T) {
	service := blog.Service{}

	request := &pb.DeletePostRequest{
		PostId: "1",
	}
	expectedResponse := &pb.DeletePostResponse{
		PostId: "1",
		Status: "deleted",
	}

	resultDelete, err := service.DeletePost(context.Background(), request)

	if err != nil {
		t.Errorf("got %s", err)
	}

	if resultDelete.Status != expectedResponse.Status {
		t.Errorf("expected %s and got %s", expectedResponse.Status, resultDelete.Status)
	}
}

func TestNotFoundBeforeDelete(t *testing.T) {
	service := blog.Service{}

	requestFilled := &pb.DeletePostRequest{
		PostId: "2",
	}

	expectedResponse := &pb.DeletePostResponse{
		PostId: "2",
		Status: "not found",
	}

	resultNotFound, err := service.DeletePost(context.Background(), requestFilled)

	if err != nil {
		t.Errorf("got %s", err)
	}

	if resultNotFound.Status != expectedResponse.Status {
		t.Errorf("expected %s and got %s", expectedResponse.Status, resultNotFound)
	}
}
