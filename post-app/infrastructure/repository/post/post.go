package post

import (
	"post-app/entity/post"
	postRepo "post-app/usecase/repository/post"
	"time"
)

type postRepositoryImpl struct {
}

func NewPostRepositoryImpl() postRepo.PostRepository {
	return &postRepositoryImpl{}
}

func (r *postRepositoryImpl) FindByID(id int) (*post.Post, error) {
	return post.NewPost(
		id,
		"dummy",
		1,
		1,
		"dummy",
		"dummy",
		time.Now(),
	), nil
}
