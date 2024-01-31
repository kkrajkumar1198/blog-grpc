package http

import (
	"net/http"

	"github.com/kkrajkumar1198/blog-grpc/internal/blog/models"
	pb "github.com/kkrajkumar1198/blog-grpc/internal/blog/protos/bin"
	"github.com/kkrajkumar1198/blog-grpc/pkg/client"
	"github.com/gin-gonic/gin"
)

// @title			HTTP Client for gRPC Client
// @version		1.0
// @description	This client handles data for sending data to gRPC client and after that to gRPC server
// @termsOfService	http://swagger.io/terms/

// @host		localhost:8080
// @BasePath	/api/v1

// getPost godoc
// @Summary      Get posts from DB
// @Description  Through a get request the id is sent to gRPC client
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Post ID"
// @Success      200  {object}  models.Post
// @Failure	 400 {object}  httpResponse
// @Failure	 404 {object} httpResponse
// @Router       /post/{id} [get]
func getPost(context *gin.Context) {
	id := context.Param("id")

	badReq := httpResponse{
		Response: "ID empty",
	}

	if id == "" {
		context.JSON(http.StatusBadRequest, badReq)
	}

	requestPB := &pb.GetPostRequest{
		Id: id,
	}

	response, err := client.ReadPost(requestPB)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	if response.Post.Id == "" {
		context.JSON(http.StatusNotFound, httpResponse{
			Response: "Not found",
		})
		return
	}

	context.JSON(http.StatusOK, response)
}

// postPost godoc
// @Summary      Creates new post
// @Description  This endpoint is for creating posts
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        post body  models.Post true "Creates post"
// @Success      200  {object}  httpResponse
// @Failure	 400 {object}  httpResponse
// @Router       /post [post]
func postPost(context *gin.Context) {
	postModel := models.Post{}
	postProtoModel := &pb.Post{}

	err := context.BindJSON(&postModel)

	postProtoModel.Title = postModel.Title
	postProtoModel.Content = postModel.Content
	postProtoModel.Author = postModel.Author
	postProtoModel.PublicationDate = postModel.PublicationDate
	postProtoModel.Tags = postModel.Tags

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := client.CreatePost(postProtoModel)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	created := httpResponse{
		Response: response,
	}

	context.JSON(http.StatusOK, created)
}

// deletePost godoc
// @Summary      Deletes post by ID
// @Description  This endpoint is for deleting post by ID
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id  path string true "uuid formatted ID"
// @Success      200  {object}  httpResponse
// @Failure	 400 {object}  httpResponse
// @Failure	 404 {object} httpResponse
// @Router       /post/{id} [delete]
func deletePost(context *gin.Context) {
	id := context.Param("id")
	requestPB := &pb.DeletePostRequest{
		Id: id,
	}

	if id == "" {
		context.JSON(http.StatusBadRequest, httpResponse{
			Response: "ID is missing",
		})
		return
	}

	response, err := client.DeletePost(requestPB)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	if response.Status == "not found" {
		context.JSON(http.StatusNotFound, httpResponse{
			Response: "Not found",
		})
		return
	}

	context.JSON(http.StatusOK, response)
}
