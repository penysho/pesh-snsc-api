package interactor

import (
	"post-app/interface/presenter"
	postRepo "post-app/usecase/repository/post"

	"github.com/gin-gonic/gin"
)

type PostInteractor interface {
	GetPost(c *gin.Context, id int, outputBoundary presenter.PostPresenter)
}

type postInteractorImpl struct {
	repository postRepo.PostRepository
}

func NewPostInteractor(postRepo postRepo.PostRepository) PostInteractor {
	return &postInteractorImpl{repository: postRepo}
}

func (i *postInteractorImpl) GetPost(c *gin.Context, id int, outputBoundary presenter.PostPresenter) {
	post, err := i.repository.FindByID(id)
	if err != nil {
		outputBoundary.ErrorResponse(err)
		return
	}
	outputBoundary.PresentGetPost(post)
}
