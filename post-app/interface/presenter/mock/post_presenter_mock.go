// Code generated by MockGen. DO NOT EDIT.
// Source: post_presenter_interface.go
//
// Generated by this command:
//
//	mockgen -source=post_presenter_interface.go -destination=mock/post_presenter_mock.go -package=presenter_mock
//

// Package presenter_mock is a generated GoMock package.
package presenter_mock

import (
	reflect "reflect"

	post "github.com/penysho/pesh-snsc-api/post-app/entity/post"
	gomock "go.uber.org/mock/gomock"
)

// MockPostPresenter is a mock of PostPresenter interface.
type MockPostPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockPostPresenterMockRecorder
}

// MockPostPresenterMockRecorder is the mock recorder for MockPostPresenter.
type MockPostPresenterMockRecorder struct {
	mock *MockPostPresenter
}

// NewMockPostPresenter creates a new mock instance.
func NewMockPostPresenter(ctrl *gomock.Controller) *MockPostPresenter {
	mock := &MockPostPresenter{ctrl: ctrl}
	mock.recorder = &MockPostPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostPresenter) EXPECT() *MockPostPresenterMockRecorder {
	return m.recorder
}

// ErrorResponse mocks base method.
func (m *MockPostPresenter) ErrorResponse(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ErrorResponse", err)
}

// ErrorResponse indicates an expected call of ErrorResponse.
func (mr *MockPostPresenterMockRecorder) ErrorResponse(err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorResponse", reflect.TypeOf((*MockPostPresenter)(nil).ErrorResponse), err)
}

// PresentGetPost mocks base method.
func (m *MockPostPresenter) PresentGetPost(post *post.Post) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PresentGetPost", post)
}

// PresentGetPost indicates an expected call of PresentGetPost.
func (mr *MockPostPresenterMockRecorder) PresentGetPost(post any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentGetPost", reflect.TypeOf((*MockPostPresenter)(nil).PresentGetPost), post)
}
