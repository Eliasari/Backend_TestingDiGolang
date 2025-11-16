package repository

import (
    "context"
    "go-fiber/app/model"
)

type IUserRepository interface {
    FindUserByUsernameOrEmail(ctx context.Context, identifier string) (*model.User, error)
    GetUsers(ctx context.Context, search string, limit, offset int64) ([]model.User, error)
    CountUsers(ctx context.Context, search string) (int64, error)
}
