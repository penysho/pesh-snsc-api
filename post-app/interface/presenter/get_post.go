package presenter

import (
	"net/http"
	server "post-app/apidef/generated"
	domainError "post-app/entity/error"
	"post-app/entity/post"

	"github.com/gin-gonic/gin"
)

type PostPresenter interface {
	PresentGetPost(post *post.Post)
	ErrorResponse(err error)
}

type postPresenterImpl struct {
	context *gin.Context
}

func NewPostPresenter(c *gin.Context) PostPresenter {
	return &postPresenterImpl{
		context: c,
	}
}

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

func (p *postPresenterImpl) ErrorResponse(err error) {
	var status int
	var message string

	switch err {
	case domainError.NotFound:
		status = http.StatusNotFound
		message = "Post not found"
	case domainError.InvalidInput:
		status = http.StatusBadRequest
		message = "Invalid input"
	default:
		status = http.StatusInternalServerError
		message = "Internal server error"
	}

	p.context.JSON(status, gin.H{"errors": message})
}
