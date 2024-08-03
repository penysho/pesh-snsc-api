package mock

import (
	"slices"

	"github.com/gin-gonic/gin"
	controller_mock "github.com/penysho/pesh-snsc-api/post-app/interface/controller/mock"
	presenter_mock "github.com/penysho/pesh-snsc-api/post-app/interface/presenter/mock"
	interactor_mock "github.com/penysho/pesh-snsc-api/post-app/usecase/interactor/mock"
	post_mock "github.com/penysho/pesh-snsc-api/post-app/usecase/repository/mock"
	"golang.org/x/exp/maps"
)

// Mock 各テストで利用されるモックを一括定義する
type Mock struct {
	OutputBoundary   *presenter_mock.MockPostPresenter
	PostRepo         *post_mock.MockPostRepository
	Interactor       *interactor_mock.MockPostInteractor
	InteractProvider *controller_mock.MockInteractProvider
}

// SetMockFuncs モック関数をセットする
// 各テストのデフォルトのモック関数群のmockFuncsと、テストケース毎に個別に利用するモック関数群のeachMockFuncsを引数に取る
// ループ処理の中でインデックスを検査し、eachMockFuncsを優先してモック関数をセットする
func SetMockFuncs(
	context *gin.Context,
	eachMocks *Mock,
	mockFuncs map[int]func(*Mock, *gin.Context),
	eachMockFuncs map[int]func(*Mock, *gin.Context),
) {
outerLoop:
	for mockIndex, mockFunc := range mockFuncs {
		eachMockFuncKeys := maps.Keys(eachMockFuncs)
		slices.Sort(eachMockFuncKeys)

		for eachMockIndex, eachMockFunc := range eachMockFuncs {
			// 同じインデックスを持つモック関数が存在すればeachMockFuncを優先してセットする
			if mockIndex == eachMockIndex {
				eachMockFunc(eachMocks, context)
				continue outerLoop
			}
			// eachMockFuncsの最後の関数をセットしたらそれ以上モック関数をセットしない
			if eachMockIndex == eachMockFuncKeys[len(eachMockFuncKeys)-1] {
				break outerLoop
			}
		}
		mockFunc(eachMocks, context)
	}
}
