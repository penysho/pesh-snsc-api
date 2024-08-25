package post_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	domainError "github.com/penysho/pesh-snsc-api/post-app/entity/error"
	"github.com/penysho/pesh-snsc-api/post-app/entity/post"
	presenter "github.com/penysho/pesh-snsc-api/post-app/interface/presenter/post"
	"github.com/stretchr/testify/assert"
)

func Test_postPresenterImpl_PresentGetPost(t *testing.T) {
	type args struct {
		post *post.Post
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "適切なステータスコードでレスポンスされる",
			args: args{
				post: post.NewPost(
					uint64(10000),
					"title",
					1,
					1,
					"caption",
					"https://example.com",
					time.Now(),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			p := presenter.NewPostPresenter(c)
			p.PresentGetPost(tt.args.post)
			assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		})
	}
}

func Test_postPresenterImpl_ErrorResponse(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name       string
		args       args
		statusCode int
		message    presenter.PostErrorMessage
	}{
		{
			name: "NotFoundが返却される",
			args: args{
				err: domainError.NotFound,
			},
			statusCode: http.StatusNotFound,
			message:    presenter.PostNotFound,
		},
		{
			name: "BadRequestが返却される",
			args: args{
				err: domainError.InvalidInput,
			},
			statusCode: http.StatusBadRequest,
			message:    presenter.InvalidInput,
		},
		{
			name: "InternalServerErrorが返却される",
			args: args{
				err: domainError.SystemError,
			},
			statusCode: http.StatusInternalServerError,
			message:    presenter.InternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			p := presenter.NewPostPresenter(c)
			p.ErrorResponse(tt.args.err)
			assert.Equal(t, tt.statusCode, w.Result().StatusCode)
			assert.Contains(t, w.Body.String(), tt.message)
		})
	}
}
