package controller

import (
	postInteractor "post-app/usecase/interactor/post"

	"github.com/gin-gonic/gin"
)

type InteractProvider interface {
	ProvidePostInteractor(c *gin.Context) postInteractor.PostInteractor
}
