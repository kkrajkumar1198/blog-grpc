package data

import (
	"github.com/kkrajkumar1198/blog-grpc/internal/database"
	"github.com/kkrajkumar1198/blog-grpc/internal/blog/models"
)

var dBConnectionObject database.IConnection

type PostDataAccess struct {
	DB   database.IConnection
	post models.Post
}

func init() {
	dBConnectionObject = database.SQLiteDB{}
}

func (p PostDataAccess) Create(post models.Post) string {
	connection, err := dBConnectionObject.GetConnection()

	if err != nil {
		return "DB ERROR: " + err.Error()
	}

	connection.Create(&post)

	return "created"
}

func (p PostDataAccess) Read(postID string) (*models.Post, error) {
	connection, err := dBConnectionObject.GetConnection()

	post := models.Post{}
	if err != nil {
		return nil, err
	}

	connection.First(&post, "post_id = ?", postID)

	return &post, nil
}

func (p PostDataAccess) Delete(post *models.Post) (*models.Post, error) {
	connection, err := dBConnectionObject.GetConnection()

	if err != nil {
		return nil, err
	}

	connection.Delete(post)

	return &models.Post{}, nil
}
