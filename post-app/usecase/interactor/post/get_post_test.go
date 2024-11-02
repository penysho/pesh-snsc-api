package post_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	domainError "github.com/penysho/pesh-snsc-api/post-app/entity/error"
	"github.com/penysho/pesh-snsc-api/post-app/entity/post"
	presenter_mock "github.com/penysho/pesh-snsc-api/post-app/interface/presenter/mock"
	interactor "github.com/penysho/pesh-snsc-api/post-app/usecase/interactor/post"
	post_mock "github.com/penysho/pesh-snsc-api/post-app/usecase/repository/mock"
	"github.com/penysho/pesh-snsc-api/post-app/util/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_postInteractorImpl_GetPost(t *testing.T) {
	const (
		FindByID = iota + 1
		PresentGetPost
	)

	mockFuncs := map[int]func(*mock.Mock, *gin.Context){
		FindByID: func(m *mock.Mock, c *gin.Context) {
			m.PostRepo.EXPECT().FindByID(gomock.Any()).Return(nil, nil)
		},
		PresentGetPost: func(m *mock.Mock, c *gin.Context) {
			m.OutputBoundary.EXPECT().PresentGetPost(gomock.Any()).DoAndReturn(
				func(_ *post.Post) {
					c.JSON(http.StatusOK, nil)
				},
			)
		},
	}

	type args struct {
		id post.ID
	}
	tests := []struct {
		name           string
		args           args
		eachMockFuncs  map[int]func(*mock.Mock, *gin.Context)
		wantStatusCode int
	}{
		{
			name: "正常系",
			args: args{
				id: 1,
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "異常系",
			args: args{
				id: 1,
			},
			eachMockFuncs: map[int]func(*mock.Mock, *gin.Context){
				FindByID: func(m *mock.Mock, c *gin.Context) {
					m.PostRepo.EXPECT().FindByID(gomock.Any()).Return(nil, domainError.NotFound)
				},
				PresentGetPost: func(m *mock.Mock, c *gin.Context) {
					m.OutputBoundary.EXPECT().ErrorResponse(gomock.Any()).DoAndReturn(
						func(_ error) {
							c.JSON(http.StatusNotFound, nil)
						},
					)
				},
			},
			wantStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			eachMocks := &mock.Mock{
				OutputBoundary: presenter_mock.NewMockPostPresenter(ctrl),
				PostRepo:       post_mock.NewMockPostRepository(ctrl),
			}

			mock.SetMockFuncs(c, eachMocks, mockFuncs, tt.eachMockFuncs)

			i := interactor.NewPostInteractor(
				eachMocks.PostRepo,
			)

			i.GetPost(c, tt.args.id, eachMocks.OutputBoundary)
			assert.Equal(t, tt.wantStatusCode, w.Result().StatusCode)
		})
	}
}
