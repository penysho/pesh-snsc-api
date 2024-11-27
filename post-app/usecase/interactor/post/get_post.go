package post

import (
	"context"

	"github.com/penysho/pesh-snsc-api/post-app/entity/post"
	"github.com/penysho/pesh-snsc-api/post-app/interface/presenter"
)

// GetPost 投稿情報取得してレスポンスする処理を操作する
func (i *postInteractorImpl) GetPost(c context.Context, id post.ID, outputBoundary presenter.PostPresenter) {
	post, err := i.repository.FindByID(id)
	if err != nil {
		outputBoundary.ErrorResponse(err)
		return
	}
	outputBoundary.PresentGetPost(post)
}
