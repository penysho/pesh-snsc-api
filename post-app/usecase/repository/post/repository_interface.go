package post

import "github.com/penysho/pesh-snsc-api/post-app/entity/post"

//go:generate mockgen -source=repository_interface.go -destination=mock/repository_interface_mock.go -package=post_mock
type PostRepository interface {
	FindByID(id int) (*post.Post, error)
}
