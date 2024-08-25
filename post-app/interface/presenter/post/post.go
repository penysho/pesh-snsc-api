package post

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/penysho/pesh-snsc-api/post-app/apidef/generated/server"
	domainError "github.com/penysho/pesh-snsc-api/post-app/entity/error"
	"github.com/penysho/pesh-snsc-api/post-app/entity/post"
	"github.com/penysho/pesh-snsc-api/post-app/interface/presenter"
)

type postPresenterImpl struct {
	context *gin.Context
}

func NewPostPresenter(c *gin.Context) presenter.PostPresenter {
	return &postPresenterImpl{
		context: c,
	}
}

// PresentGetPost 成功レスポンスを返却する
func (p *postPresenterImpl) PresentGetPost(post *post.Post) {
	p.context.JSON(http.StatusOK, &server.GetPostResponse{
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
}

type PostErrorMessage = string

const (
	PostNotFound        PostErrorMessage = "Post not found"
	InvalidInput        PostErrorMessage = "Invalid input"
	InternalServerError PostErrorMessage = "Internal server error"
)

// ErrorResponse エラーレスポンスを返却する
func (p *postPresenterImpl) ErrorResponse(err error) {
	var status int
	var message PostErrorMessage

	switch err {
	case domainError.NotFound:
		status = http.StatusNotFound
		message = PostNotFound
	case domainError.InvalidInput:
		status = http.StatusBadRequest
		message = InvalidInput
	default:
		status = http.StatusInternalServerError
		message = InternalServerError
	}

	p.context.JSON(status, server.DomainError{
		Message: message,
	})
}
