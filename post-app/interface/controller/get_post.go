package controller

import (
	"post-app/interface/presenter"

	"github.com/gin-gonic/gin"
)

// GetPost 投稿詳細情報の取得
// (GET /post/{postId})
func (s *Server) GetPost(c *gin.Context, postId int) {
	postInteractor := s.InteractProvider.ProvidePostInteractor(c)
	presenter := presenter.NewPostPresenter(c)
	postInteractor.GetPost(c, postId, presenter)
}
