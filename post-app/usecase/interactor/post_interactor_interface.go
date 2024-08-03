package interactor

import (
	"github.com/gin-gonic/gin"
	"github.com/penysho/pesh-snsc-api/post-app/interface/presenter"
)

//go:generate mockgen -source=post_interactor_interface.go -destination=mock/post_interactor_mock.go -package=interactor_mock
type PostInteractor interface {
	GetPost(c *gin.Context, id uint64, outputBoundary presenter.PostPresenter)
}
