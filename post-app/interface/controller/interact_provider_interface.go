package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/penysho/pesh-snsc-api/post-app/usecase/interactor"
)

//go:generate mockgen -source=interact_provider_interface.go -destination=mock/interact_provider_mock.go -package=controller_mock
type InteractProvider interface {
	ProvidePostInteractor(c *gin.Context) interactor.PostInteractor
}
