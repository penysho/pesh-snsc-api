// Code generated by MockGen. DO NOT EDIT.
// Source: repository_interface.go
//
// Generated by this command:
//
//	mockgen -source=repository_interface.go -destination=mock/repository_interface_mock.go -package=post_mock
//

// Package post_mock is a generated GoMock package.
package post_mock

import (
	reflect "reflect"

	post "github.com/penysho/pesh-snsc-api/post-app/entity/post"
	gomock "go.uber.org/mock/gomock"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockPostRepository) FindByID(id int) (*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockPostRepositoryMockRecorder) FindByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPostRepository)(nil).FindByID), id)
}
