package post

import (
	"github.com/penysho/pesh-snsc-api/post-app/usecase/interactor"
	"github.com/penysho/pesh-snsc-api/post-app/usecase/repository"
)

type postInteractorImpl struct {
	repository repository.PostRepository
}

func NewPostInteractor(postRepo repository.PostRepository) interactor.PostInteractor {
	return &postInteractorImpl{repository: postRepo}
}
