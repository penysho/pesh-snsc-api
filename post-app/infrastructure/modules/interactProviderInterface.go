package modules

import (
	"post-app/interface/controller"
	postInteractor "post-app/usecase/interactor/post"

	"github.com/gin-gonic/gin"
)

type interactProviderImpl struct{}

func NewInteractProvider() controller.InteractProvider {
	return &interactProviderImpl{}
}

func (p *interactProviderImpl) ProvidePostInteractor(c *gin.Context) postInteractor.PostInteractor {
	return c.MustGet("postInteractor").(postInteractor.PostInteractor)
}
