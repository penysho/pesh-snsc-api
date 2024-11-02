package repository

import "github.com/penysho/pesh-snsc-api/post-app/entity/post"

//go:generate mockgen -source=post_repository_interface.go -destination=mock/post_repository_mock.go -package=repository_mock
type PostRepository interface {
	FindByID(id post.ID) (*post.Post, error)
}
