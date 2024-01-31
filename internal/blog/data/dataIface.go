package data

import (
	"github.com/kkrajkumar1198/blog-grpc/internal/blog/models"
)

var dataAccessIface IPostDataAccess

type IPostDataAccess interface {
	Create(post models.Post) string
	Read(postID string) (*models.Post, error)
	Delete(postIDOrPost interface{}) (*models.Post, error)
}
