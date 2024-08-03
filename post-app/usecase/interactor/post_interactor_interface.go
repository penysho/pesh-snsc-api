package interactor

import (
	"context"

	"github.com/penysho/pesh-snsc-api/post-app/interface/presenter"
)

//go:generate mockgen -source=post_interactor_interface.go -destination=mock/post_interactor_mock.go -package=interactor_mock
type PostInteractor interface {
	GetPost(c context.Context, id uint64, outputBoundary presenter.PostPresenter)
}
