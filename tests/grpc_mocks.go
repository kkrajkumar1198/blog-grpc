package tests

import "github.com/kkrajkumar1198/blog-grpc/internal/blog/models"

type mockDataAccess struct {
}

func (m mockDataAccess) Create(post models.Post) string {
	return "created"
}

func (m mockDataAccess) Read(id string) (*models.Post, error) {
	if id == "2" {
		return &models.Post{
			PostID: "",
		}, nil
	}
	return &models.Post{
		PostID:          "1",
		Title:           "Sample Title",
		Content:         "Sample Content",
		Author:          "Sample Author",
		PublicationDate: "2024-01-31",
		Tags:            "tag1 ,tag2",
	}, nil
}

func (m mockDataAccess) Delete(postIDOrPost interface{}) (*models.Post, error) {
	return &models.Post{}, nil
}
