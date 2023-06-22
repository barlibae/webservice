// Code generated by MockGen. DO NOT EDIT.
// Source: albums.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	albums "github.com/barlibae/webservice/albums"
	gomock "github.com/golang/mock/gomock"
)

// MockAlbumRepo is a mock of AlbumRepo interface.
type MockAlbumRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumRepoMockRecorder
}

// MockAlbumRepoMockRecorder is the mock recorder for MockAlbumRepo.
type MockAlbumRepoMockRecorder struct {
	mock *MockAlbumRepo
}

// NewMockAlbumRepo creates a new mock instance.
func NewMockAlbumRepo(ctrl *gomock.Controller) *MockAlbumRepo {
	mock := &MockAlbumRepo{ctrl: ctrl}
	mock.recorder = &MockAlbumRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlbumRepo) EXPECT() *MockAlbumRepoMockRecorder {
	return m.recorder
}

// AlbumByID mocks base method.
func (m *MockAlbumRepo) AlbumByID(id string) (albums.Album, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlbumByID", id)
	ret0, _ := ret[0].(albums.Album)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// AlbumByID indicates an expected call of AlbumByID.
func (mr *MockAlbumRepoMockRecorder) AlbumByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlbumByID", reflect.TypeOf((*MockAlbumRepo)(nil).AlbumByID), id)
}

// Albums mocks base method.
func (m *MockAlbumRepo) Albums() []albums.Album {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Albums")
	ret0, _ := ret[0].([]albums.Album)
	return ret0
}

// Albums indicates an expected call of Albums.
func (mr *MockAlbumRepoMockRecorder) Albums() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Albums", reflect.TypeOf((*MockAlbumRepo)(nil).Albums))
}

// CreateAlbum mocks base method.
func (m *MockAlbumRepo) CreateAlbum(album albums.Album) albums.Album {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlbum", album)
	ret0, _ := ret[0].(albums.Album)
	return ret0
}

// CreateAlbum indicates an expected call of CreateAlbum.
func (mr *MockAlbumRepoMockRecorder) CreateAlbum(album interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlbum", reflect.TypeOf((*MockAlbumRepo)(nil).CreateAlbum), album)
}