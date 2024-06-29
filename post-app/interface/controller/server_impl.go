package controller

import (
	postInteractor "post-app/usecase/interactor/post"

	"github.com/gin-gonic/gin"
)

type Server struct {
	InteractProvider
}

func NewServer(interactProvider InteractProvider) *Server {
	return &Server{
		InteractProvider: interactProvider,
	}
}

type InteractProvider interface {
	ProvidePostInteractor(c *gin.Context) postInteractor.PostInteractor
}

type interactProviderImpl struct{}

func NewInteractProvider() InteractProvider {
	return &interactProviderImpl{}
}

func (p *interactProviderImpl) ProvidePostInteractor(c *gin.Context) postInteractor.PostInteractor {
	return c.MustGet("postInteractor").(postInteractor.PostInteractor)
}
