package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/penysho/pesh-snsc-api/post-app/interface/presenter"
)

// GetPost 投稿詳細情報の取得
// (GET /post-app/post/{postId})
func (s *Server) GetPost(c *gin.Context, postId int) {
	postInteractor := s.InteractProvider.ProvidePostInteractor(c)
	presenter := presenter.NewPostPresenter(c)
	postInteractor.GetPost(c, postId, presenter)
}
