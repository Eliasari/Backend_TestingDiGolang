package service_test

import (
	"context"
	"errors"
	"testing"

	"go-fiber/app/model"
	"go-fiber/app/repository"
	"go-fiber/app/service"
	"go-fiber/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginServiceMongo_Success(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(repository.MockUserRepository)

	hashed, _ := utils.HashPassword("123") 

	fakeUser := &model.User{
		Username: "elia",
		Password: hashed,
		Role:     "user",
	}

	mockRepo.
		On("FindUserByUsernameOrEmail", mock.Anything, "elia").
		Return(fakeUser, nil)

	req := model.LoginRequest{
		Username: "elia",
		Password: "123",
	}

	resp, err := service.LoginServiceMongo(ctx, mockRepo, req)

	assert.NoError(t, err)
	assert.Equal(t, fakeUser.Username, resp.User.Username)
	assert.NotEmpty(t, resp.Token)

}

func TestLoginServiceMongo_UserNotFound(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)

	mockRepo.
		On("FindUserByUsernameOrEmail", mock.Anything, "unknown").
		Return(nil, errors.New("not found"))

	req := model.LoginRequest{
		Username: "unknown",
		Password: "wrong",
	}

	ctx := context.Background()

	resp, err := service.LoginServiceMongo(ctx, mockRepo, req)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "username atau password salah")
}

func TestLoginServiceMongo_InvalidPassword(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)

	hashed, _ := utils.HashPassword("correctpass")

	user := &model.User{
		Username: "testuser",
		Password: hashed,
	}

	mockRepo.
		On("FindUserByUsernameOrEmail", mock.Anything, "testuser").
		Return(user, nil)

	req := model.LoginRequest{
		Username: "testuser",
		Password: "wrongpass",
	}

	ctx := context.Background()

	resp, err := service.LoginServiceMongo(ctx, mockRepo, req)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "username atau password salah")
}

func TestGetUsers_Success(t *testing.T) {
    mockRepo := new(repository.MockUserRepository)
    ctx := context.Background()

    // mock data
    users := []model.User{
        {Username: "elia", Email: "elia@test.com"},
        {Username: "akmal", Email: "akmal@test.com"},
    }

    search := ""
    limit := int64(10)
    offset := int64(0)

    // Mock expectation
    mockRepo.On("GetUsers", ctx, search, limit, offset).
        Return(users, nil)

    mockRepo.On("CountUsers", ctx, search).
        Return(int64(len(users)), nil)

    // ACT
    gotUsers, err := mockRepo.GetUsers(ctx, search, limit, offset)
    total, errCount := mockRepo.CountUsers(ctx, search)

    // ASSERT
    assert.Nil(t, err)
    assert.Nil(t, errCount)
    assert.Equal(t, 2, len(gotUsers))
    assert.Equal(t, int64(2), total)

    mockRepo.AssertExpectations(t)
}

func TestGetUsers_EmptyList(t *testing.T) {
    mockRepo := new(repository.MockUserRepository)
    ctx := context.Background()

    search := "zzz"
    limit := int64(10)
    offset := int64(0)

    mockRepo.On("GetUsers", ctx, search, limit, offset).
        Return([]model.User{}, nil)

    mockRepo.On("CountUsers", ctx, search).
        Return(int64(0), nil)

    gotUsers, err := mockRepo.GetUsers(ctx, search, limit, offset)
    total, errCount := mockRepo.CountUsers(ctx, search)

    assert.Nil(t, err)
    assert.Nil(t, errCount)
    assert.Equal(t, 0, len(gotUsers))
    assert.Equal(t, int64(0), total)

    mockRepo.AssertExpectations(t)
}

func TestGetUsers_Error(t *testing.T) {
    mockRepo := new(repository.MockUserRepository)
    ctx := context.Background()

    search := ""
    limit := int64(10)
    offset := int64(0)

    // mock return nil slice
    mockRepo.On("GetUsers", ctx, search, limit, offset).
        Return([]model.User(nil), assert.AnError)

    gotUsers, err := mockRepo.GetUsers(ctx, search, limit, offset)

    assert.Nil(t, gotUsers)      // NOW VALID
    assert.Error(t, err)

    mockRepo.AssertExpectations(t)
}
