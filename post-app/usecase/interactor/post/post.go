package interactor

import (
	"post-app/interface/presenter"
	postRepo "post-app/usecase/repository/post"

	"github.com/gin-gonic/gin"
)

type PostInteractor struct {
	repository postRepo.PostRepository
}

func NewPostInteractor(postRepo postRepo.PostRepository) *PostInteractor {
	return &PostInteractor{repository: postRepo}
}

func (s *PostInteractor) GetPost(c *gin.Context, id int, outputBoundary presenter.PostPresenter) error {
	post, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}
	return outputBoundary.PresentGetPost(post)
}
