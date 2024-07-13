package interactor

import (
	postRepo "github.com/penysho/pesh-snsc-api/post-app/usecase/repository/post"

	"github.com/gin-gonic/gin"
	"github.com/penysho/pesh-snsc-api/post-app/interface/presenter"
)

//go:generate mockgen -source=post.go -destination=mock/post_mock.go -package=interactor_mock
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
