// Code generated by MockGen. DO NOT EDIT.
// Source: post_interactor_interface.go
//
// Generated by this command:
//
//	mockgen -source=post_interactor_interface.go -destination=mock/post_interactor_mock.go -package=interactor_mock
//

// Package interactor_mock is a generated GoMock package.
package interactor_mock

import (
	context "context"
	reflect "reflect"

	post "github.com/penysho/pesh-snsc-api/post-app/entity/post"
	presenter "github.com/penysho/pesh-snsc-api/post-app/interface/presenter"
	gomock "go.uber.org/mock/gomock"
)

// MockPostInteractor is a mock of PostInteractor interface.
type MockPostInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockPostInteractorMockRecorder
	isgomock struct{}
}

// MockPostInteractorMockRecorder is the mock recorder for MockPostInteractor.
type MockPostInteractorMockRecorder struct {
	mock *MockPostInteractor
}

// NewMockPostInteractor creates a new mock instance.
func NewMockPostInteractor(ctrl *gomock.Controller) *MockPostInteractor {
	mock := &MockPostInteractor{ctrl: ctrl}
	mock.recorder = &MockPostInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostInteractor) EXPECT() *MockPostInteractorMockRecorder {
	return m.recorder
}

// GetPost mocks base method.
func (m *MockPostInteractor) GetPost(c context.Context, id post.ID, outputBoundary presenter.PostPresenter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetPost", c, id, outputBoundary)
}

// GetPost indicates an expected call of GetPost.
func (mr *MockPostInteractorMockRecorder) GetPost(c, id, outputBoundary any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockPostInteractor)(nil).GetPost), c, id, outputBoundary)
}
