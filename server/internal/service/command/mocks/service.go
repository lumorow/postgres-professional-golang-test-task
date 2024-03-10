// Code generated by MockGen. DO NOT EDIT.
// Source: command_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	entity "pstgrprof/server/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateCommand mocks base method.
func (m *MockRepository) CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommand", ctx, command)
	ret0, _ := ret[0].(*entity.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCommand indicates an expected call of CreateCommand.
func (mr *MockRepositoryMockRecorder) CreateCommand(ctx, command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommand", reflect.TypeOf((*MockRepository)(nil).CreateCommand), ctx, command)
}

// GetAllCommands mocks base method.
func (m *MockRepository) GetAllCommands(ctx context.Context) (*[]entity.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCommands", ctx)
	ret0, _ := ret[0].(*[]entity.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCommands indicates an expected call of GetAllCommands.
func (mr *MockRepositoryMockRecorder) GetAllCommands(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCommands", reflect.TypeOf((*MockRepository)(nil).GetAllCommands), ctx)
}

// GetCommandById mocks base method.
func (m *MockRepository) GetCommandById(ctx context.Context, id int64) (*entity.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommandById", ctx, id)
	ret0, _ := ret[0].(*entity.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommandById indicates an expected call of GetCommandById.
func (mr *MockRepositoryMockRecorder) GetCommandById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommandById", reflect.TypeOf((*MockRepository)(nil).GetCommandById), ctx, id)
}

// GetCommands mocks base method.
func (m *MockRepository) GetCommands(ctx context.Context, id []int64) (*[]entity.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommands", ctx, id)
	ret0, _ := ret[0].(*[]entity.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommands indicates an expected call of GetCommands.
func (mr *MockRepositoryMockRecorder) GetCommands(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommands", reflect.TypeOf((*MockRepository)(nil).GetCommands), ctx, id)
}

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// CheckKey mocks base method.
func (m *MockCache) CheckKey(key int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckKey", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckKey indicates an expected call of CheckKey.
func (mr *MockCacheMockRecorder) CheckKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckKey", reflect.TypeOf((*MockCache)(nil).CheckKey), key)
}

// Delete mocks base method.
func (m *MockCache) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCache)(nil).Delete), key)
}

// GetAll mocks base method.
func (m *MockCache) GetAll() ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCacheMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCache)(nil).GetAll))
}

// Set mocks base method.
func (m *MockCache) Set(key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCache)(nil).Set), key, value)
}
