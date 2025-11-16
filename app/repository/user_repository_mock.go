package repository

import (
    "context"
    "go-fiber/app/model"
    "github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) FindUserByUsernameOrEmail(ctx context.Context, identifier string) (*model.User, error) {
    args := m.Called(ctx, identifier)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetUsers(ctx context.Context, search string, limit, offset int64) ([]model.User, error) {
    args := m.Called(ctx, search, limit, offset)
    return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) CountUsers(ctx context.Context, search string) (int64, error) {
    args := m.Called(ctx, search)
    return args.Get(0).(int64), args.Error(1)
}
