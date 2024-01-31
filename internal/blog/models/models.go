// entities/post.go

package models

type Post struct {
	PostID          string `json:"post_id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	Author          string `json:"author"`
	PublicationDate string `json:"publication_date"`
	Tags            string `gorm:"type:TEXT"`
}
