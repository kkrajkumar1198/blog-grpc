// post-repo.go

package repos

import (
	"fmt"
	"os"

	"github.com/kkrajkumar1198/blog-grpc/entities"
	"github.com/kkrajkumar1198/blog-grpc/protobuf/protobuf_blog"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamp"
)

const STORAGE_FILE = "./posts-storage.pb"

type PostRepo struct {
	posts []entities.Post
}

func NewPostRepo() *PostRepo {
	var pr = PostRepo{make([]entities.Post, 0)}
	pr.loadFromFileStorage()
	return &pr
}

func (p *PostRepo) Create(post entities.Post) entities.Post {
	newPost := entities.Post{
		PostID:          fmt.Sprint(len(p.posts) + 1),
		Title:           post.Title,
		Content:         post.Content,
		Author:          post.Author,
		PublicationDate: post.PublicationDate,
		Tags:            post.Tags,
	}

	p.posts = append(p.posts, newPost)
	p.saveToFileStorage()
	return newPost
}

func (p *PostRepo) GetList() []entities.Post {
	return p.posts
}

func (p *PostRepo) GetOne(id uint) (entities.Post, error) {
	for _, post := range p.posts {
		if post.PostID == fmt.Sprint(id) {
			return post, nil
		}
	}
	return entities.Post{}, fmt.Errorf("post with ID '%d' not found", id)
}

func (p *PostRepo) Update(id uint, updatedPost entities.Post) (entities.Post, error) {
	for i, post := range p.posts {
		if post.PostID == fmt.Sprint(id) {
			updatedPost.PostID = fmt.Sprint(id)
			p.posts = append(p.posts[:i], p.posts[i+1:]...)
			p.posts = append(p.posts, updatedPost)
			p.saveToFileStorage()
			return updatedPost, nil
		}
	}
	return entities.Post{}, fmt.Errorf("post with ID '%d' not found", id)
}

func (p *PostRepo) DeleteOne(id uint) (bool, error) {
	for i, post := range p.posts {
		if post.PostID == fmt.Sprint(id) {
			p.posts = append(p.posts[:i], p.posts[i+1:]...)
			p.saveToFileStorage()
			return true, nil
		}
	}
	return false, fmt.Errorf("post with ID '%d' not found", id)
}

func (p *PostRepo) saveToFileStorage() error {
	postsMessage := &protobuf_blog.Posts{
		Posts: make([]*protobuf_blog.Post, len(p.posts)),
	}

	for i, post := range p.posts {
		protoPost := ToProtoPost(post)
		postsMessage.Posts[i] = protoPost
	}

	data, err := proto.Marshal(postsMessage)
	if err != nil {
		return fmt.Errorf("cannot marshal to binary: %w", err)
	}

	err = os.WriteFile(STORAGE_FILE, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

func (p *PostRepo) loadFromFileStorage() error {
	_, err := os.Stat(STORAGE_FILE)
	if err != nil {
		fmt.Println("storage file is not found, starting with an empty storage")
		return nil
	}

	data, err := os.ReadFile(STORAGE_FILE)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	var postsMessage protobuf_blog.Posts
	err = proto.Unmarshal(data, &postsMessage)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary data to protobuf: %w", err)
	}

	for _, protoPost := range postsMessage.Posts {
		post := ToPost(protoPost)
		p.posts = append(p.posts, post)
	}

	return nil
}

func ToProtoPost(post entities.Post) *protobuf_blog.Post {
	return &protobuf_blog.Post{
		PostId:          post.PostID,
		Title:           post.Title,
		Content:         post.Content,
		Author:          post.Author,
		PublicationDate: post.PublicationDate,
		Tags:            post.Tags,
	}
}

func ToPost(protoPost *protobuf_blog.Post) entities.Post {
	return entities.Post{
		PostID:          protoPost.GetPostId(),
		Title:           protoPost.Title,
		Content:         protoPost.Content,
		Author:          protoPost.Author,
		PublicationDate: protoPost.PublicationDate,
		Tags:            protoPost.Tags,
	}
}
