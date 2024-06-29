package post

import (
	post "post-app/entity/post"
)

type PostRepository interface {
	FindByID(id int) (*post.Post, error)
}
