package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/penysho/pesh-snsc-api/post-app/interface/controller"
	controller_mock "github.com/penysho/pesh-snsc-api/post-app/interface/controller/mock"
	interactor_mock "github.com/penysho/pesh-snsc-api/post-app/usecase/interactor/mock"
	"github.com/penysho/pesh-snsc-api/post-app/util/mock"
	"go.uber.org/mock/gomock"
)

func TestServer_GetPost(t *testing.T) {
	const (
		ProvidePostInteractor = iota + 1
		GetPost
	)

	mockFuncs := map[int]func(*mock.Mock, *gin.Context){
		ProvidePostInteractor: func(m *mock.Mock, c *gin.Context) {
			m.InteractProvider.EXPECT().ProvidePostInteractor(gomock.Any()).Return(m.Interactor)
		},
		GetPost: func(m *mock.Mock, c *gin.Context) {
			m.Interactor.EXPECT().GetPost(gomock.Any(), gomock.Any(), gomock.Any())
		},
	}

	type args struct {
		postId uint64
	}
	tests := []struct {
		name          string
		args          args
		eachMockFuncs map[int]func(*mock.Mock, *gin.Context)
	}{
		{
			name: "正常系",
			args: args{
				postId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			eachMocks := &mock.Mock{
				Interactor:       interactor_mock.NewMockPostInteractor(ctrl),
				InteractProvider: controller_mock.NewMockInteractProvider(ctrl),
			}

			mock.SetMockFuncs(c, eachMocks, mockFuncs, tt.eachMockFuncs)

			s := controller.NewServer(eachMocks.InteractProvider)
			s.GetPost(c, tt.args.postId)
		})
	}
}
