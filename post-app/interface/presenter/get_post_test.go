package presenter_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	domainError "github.com/penysho/pesh-snsc-api/post-app/entity/error"
	"github.com/penysho/pesh-snsc-api/post-app/entity/post"
	"github.com/penysho/pesh-snsc-api/post-app/interface/presenter"
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
					int(10000),
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
		name           string
		args           args
		wantStatusCode int
	}{
		{
			name: "NotFoundが返却される",
			args: args{
				err: domainError.NotFound,
			},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "BadRequestが返却される",
			args: args{
				err: domainError.InvalidInput,
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "InternalServerErrorが返却される",
			args: args{
				err: domainError.SystemError,
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			p := presenter.NewPostPresenter(c)
			p.ErrorResponse(tt.args.err)
			assert.Equal(t, tt.wantStatusCode, w.Result().StatusCode)
		})
	}
}
