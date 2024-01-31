
package blog

import (
	"context"

	"github.com/kkrajkumar1198/blog-grpc/internal/blog/data"
	"github.com/kkrajkumar1198/blog-grpc/internal/blog/models"
	pb "github.com/kkrajkumar1198/blog-grpc/internal/blog/protos/bin"
	"github.com/google/uuid"
)

var DataAccess data.IPostDataAccess

func init() {
	DataAccess = data.PostDataAccess{}
}

type Service struct {
	pb.UnimplementedBlogServiceServer
}

func (s *Service) Create(ctx context.Context, post *pb.Post) (*pb.CreatePostResponse, error) {
	postID := uuid.New().String()

	newPost := &models.Post{
		PostID:          postID,
		Title:           post.Title,
		Content:         post.Content,
		Author:          post.Author,
		PublicationDate: post.PublicationDate,
		Tags:            post.Tags,
	}

	status := DataAccess.Create(*newPost)

	return &pb.CreatePostResponse{PostId: newPost.PostID, Response: status}, nil
}

func (s *Service) Get(ctx context.Context, request *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	postID := request.PostId
	dbResponse, err := DataAccess.Read(postID)

	if err != nil {
		return nil, err
	}

	post := &pb.Post{
		PostId:          dbResponse.PostID,
		Title:           dbResponse.Title,
		Content:         dbResponse.Content,
		Author:          dbResponse.Author,
		PublicationDate: dbResponse.PublicationDate,
		Tags:            dbResponse.Tags,
	}

	return &pb.GetPostResponse{
		Post: post,
	}, nil
}

func (s *Service) Delete(ctx context.Context, request *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	postID := request.PostId
	queryPost, err := DataAccess.Read(postID)

	if err != nil {
		return nil, err
	}

	if queryPost.PostID == "" {
		return &pb.DeletePostResponse{
			PostId: postID,
			Status: "not found",
		}, nil
	}

	DataAccess.Delete(queryPost)

	delPostResponse := &pb.DeletePostResponse{
		PostId: postID,
		Status: "deleted",
	}

	return delPostResponse, nil
}
