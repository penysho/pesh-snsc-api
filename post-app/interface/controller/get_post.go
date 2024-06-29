package controller

import (
	"net/http"
	"post-app/interface/presenter"
	postInteractor "post-app/usecase/interactor/post"
	postRepo "post-app/usecase/repository/post"

	"github.com/gin-gonic/gin"
)

// GetPost 投稿詳細情報の取得
// (GET /post/{postId})
func (s *Server) GetPost(c *gin.Context, postId int) {
	postRepo := c.MustGet("postRepo").(postRepo.PostRepository)
	presenter := presenter.NewPostPresenter(c)
	postInteractor := postInteractor.NewPostInteractor(postRepo)
	err := postInteractor.GetPost(c, postId, presenter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
}
