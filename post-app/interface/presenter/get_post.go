package presenter

import (
	"net/http"
	server "post-app/apidef/generated"
	"post-app/entity/post"

	"github.com/gin-gonic/gin"
)

type PostPresenter interface {
	PresentGetPost(post *post.Post) error
}

type postPresenterImpl struct {
	context *gin.Context
}

func NewPostPresenter(c *gin.Context) PostPresenter {
	return &postPresenterImpl{
		context: c,
	}
}

func (pp *postPresenterImpl) PresentGetPost(post *post.Post) error {
	pp.context.JSON(http.StatusOK, &server.GetPostResponse{
		Post: server.Post{
			Id:            post.GetId(),
			Title:         post.GetTitle(),
			LikeCount:     post.GetLikeCount(),
			CommentsCount: post.GetCommentsCount(),
			Caption:       post.GetCaption(),
			Permalink:     post.GetPermalink(),
			PostedAt:      post.GetPostedAt(),
		},
	})
	return nil
}
