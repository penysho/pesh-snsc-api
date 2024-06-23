package controller

import (
	"net/http"
	"post-app/interface/presenter"

	"github.com/gin-gonic/gin"
)

// GetPost 投稿詳細情報の取得
// (GET /post/{postId})
func (s *Server) GetPost(c *gin.Context, postId int) {
	presenter := presenter.NewPostPresenter(c)
	err := s.postInteractor.GetPost(postId, presenter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
}
