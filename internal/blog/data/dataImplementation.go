package data

import (
	"errors"

	"github.com/kkrajkumar1198/blog-grpc/internal/blog/models"
	database "github.com/kkrajkumar1198/blog-grpc/internal/databases"
)

var dBConnectionObject database.IConnection

type PostDataAccess struct {
	DB database.IConnection
}

func init() {
	dBConnectionObject = database.SQLiteDB{}
}

func (p *PostDataAccess) Create(post models.Post) string {
	connection, err := dBConnectionObject.GetConnection()

	if err != nil {
		return "DB ERROR: " + err.Error()
	}

	connection.Create(&post)

	return "created"
}

func (p *PostDataAccess) Read(postID string) (*models.Post, error) {
	connection, err := dBConnectionObject.GetConnection()

	post := models.Post{}
	if err != nil {
		return nil, err
	}

	connection.First(&post, "post_id = ?", postID)

	return &post, nil
}

func (p *PostDataAccess) Delete(postIDOrPost interface{}) (*models.Post, error) {
	connection, err := dBConnectionObject.GetConnection()
	var post models.Post
	if err != nil {
		return nil, err
	}

	switch v := postIDOrPost.(type) {
	case string:
		// Assuming postIDOrPost is the post ID
		connection.Where("post_id = ?", v).Delete(&post)
	case *models.Post:
		// Assuming postIDOrPost is a *models.Post
		connection.Where("post_id = ?", v.PostID).Delete(&post)
	default:
		return nil, errors.New("unsupported type for postIDOrPost")
	}

	return &models.Post{}, nil
}
