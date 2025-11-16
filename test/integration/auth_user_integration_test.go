// file: test/integration/auth_user_integration_test.go
package integration

import (
    "context"
    "testing"

    "go-fiber/app/model"
    "go-fiber/app/repository"
    "go-fiber/app/service"

    "github.com/stretchr/testify/assert"
)

func TestLoginServiceMongo_Integration_Dynamic(t *testing.T) {
    SetupEnvTest(t)

    client, coll := SetupMongoTestCollection(t, "users")
    defer CleanupMongo(t, client)

    ctx := context.Background()

    // seeding user test (hanya jika belum ada)
    SeedUsers(ctx, coll)

    repo := repository.NewUserRepository(coll.Database())

    // ambil semua user yang ada di collection
    users, err := repo.GetUsers(ctx, "", 100, 0)
    if err != nil {
        t.Fatalf("failed to get users: %v", err)
    }

    if len(users) == 0 {
        t.Fatal("no users found in DB to test login")
    }

    // test login untuk setiap user
    for _, u := range users {
        t.Run("Login "+u.Username, func(t *testing.T) {
            req := model.LoginRequest{
                Username: u.Username,
                Password: "123", // sesuai password yang diset di SeedUsers
            }
            resp, err := service.LoginServiceMongo(ctx, repo, req)
            if err != nil {
                t.Errorf("login failed for user %s: %v", u.Username, err)
                return
            }

            assert.Equal(t, u.Username, resp.User.Username)
            assert.Equal(t, u.Role, resp.User.Role)
        })
    }
}

func TestLoginServiceMongo_Integration_FullDynamic(t *testing.T) {
    SetupEnvTest(t)

    client, coll := SetupMongoTestCollection(t, "users")
    defer CleanupMongo(t, client)

    ctx := context.Background()

    // pastikan ada user test
    SeedUsers(ctx, coll)

    repo := repository.NewUserRepository(coll.Database())

    // ambil semua user yang ada di DB
    users, err := repo.GetUsers(ctx, "", 100, 0)
    if err != nil {
        t.Fatalf("failed to get users: %v", err)
    }

    if len(users) == 0 {
        t.Fatal("no users found in DB to test login")
    }

    for _, u := range users {
        t.Run("LoginSuccess_"+u.Username, func(t *testing.T) {
            req := model.LoginRequest{
                Username: u.Username,
                Password: "123", // sesuai seed
            }
            resp, err := service.LoginServiceMongo(ctx, repo, req)
            assert.NoError(t, err)
            assert.Equal(t, u.Username, resp.User.Username)
            assert.Equal(t, u.Role, resp.User.Role)
        })

        t.Run("LoginWrongPassword_"+u.Username, func(t *testing.T) {
            req := model.LoginRequest{
                Username: u.Username,
                Password: "wrongpassword",
            }
            resp, err := service.LoginServiceMongo(ctx, repo, req)
            assert.Error(t, err)
            assert.Nil(t, resp)
        })
    }

    // test login untuk username yang tidak ada di DB
    t.Run("LoginNotFound", func(t *testing.T) {
        req := model.LoginRequest{
            Username: "nonexistent_user_12345",
            Password: "123",
        }
        resp, err := service.LoginServiceMongo(ctx, repo, req)
        assert.Error(t, err)
        assert.Nil(t, resp)
    })
}

func TestGetUsers_Integration(t *testing.T) {
    SetupEnvTest(t)
    client, coll := SetupMongoTestCollection(t, "users_test")
    defer CleanupMongo(t, client)

    ctx := context.Background()
    SeedUsers(ctx, coll)

    repo := repository.NewUserRepository(coll.Database())
    users, err := repo.GetUsers(ctx, "", 10, 0)
    assert.NoError(t, err)
    assert.GreaterOrEqual(t, len(users), 2) 

    count, err := repo.CountUsers(ctx, "")
    assert.NoError(t, err)
    assert.GreaterOrEqual(t, count, int64(2))
}
