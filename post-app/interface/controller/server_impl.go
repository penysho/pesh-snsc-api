package controller

import (
	postRepoImpl "post-app/infrastructure/db/repository/post"
	postInteractor "post-app/usecase/interactor/post"
)

type Server struct {
	postInteractor *postInteractor.PostInteractor
}

func NewServer() *Server {
	postRepoImpl := postRepoImpl.NewPostRepositoryImpl()
	postInteractor := postInteractor.NewPostInteractor(postRepoImpl)
	return &Server{
		postInteractor: postInteractor,
	}
}
