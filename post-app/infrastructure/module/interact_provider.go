package module

import (
	"github.com/penysho/pesh-snsc-api/post-app/interface/controller"
	postInteractor "github.com/penysho/pesh-snsc-api/post-app/usecase/interactor/post"

	"github.com/gin-gonic/gin"
)

type interactProviderImpl struct{}

func NewInteractProvider() controller.InteractProvider {
	return &interactProviderImpl{}
}

func (p *interactProviderImpl) ProvidePostInteractor(c *gin.Context) postInteractor.PostInteractor {
	return c.MustGet("postInteractor").(postInteractor.PostInteractor)
}
