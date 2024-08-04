package presenter

import (
	"github.com/penysho/pesh-snsc-api/post-app/entity/post"
)

//go:generate mockgen -source=post_presenter_interface.go -destination=mock/post_presenter_mock.go -package=presenter_mock
type PostPresenter interface {
	PresentGetPost(post *post.Post)
	ErrorResponse(err error)
}
