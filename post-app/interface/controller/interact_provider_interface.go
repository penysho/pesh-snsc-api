package controller

import (
	postInteractor "post-app/usecase/interactor/post"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=interact_provider_interface.go -destination=mock/interact_provider_interface_mock.go -package=controller_mock
type InteractProvider interface {
	ProvidePostInteractor(c *gin.Context) postInteractor.PostInteractor
}
