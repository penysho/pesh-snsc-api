package module

import (
	"github.com/penysho/pesh-snsc-api/post-app/interface/controller"
	"github.com/penysho/pesh-snsc-api/post-app/usecase/interactor"

	"github.com/gin-gonic/gin"
)

type interactProviderImpl struct{}

func NewInteractProvider() controller.InteractProvider {
	return &interactProviderImpl{}
}

// ProvidePostInteractor PostInteractorを提供する
func (p *interactProviderImpl) ProvidePostInteractor(c *gin.Context) interactor.PostInteractor {
	return c.MustGet("postInteractor").(interactor.PostInteractor)
}
